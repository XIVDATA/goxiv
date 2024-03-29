package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model/character"
	"github.com/xivdata/goxiv/model/gear"
)

const (
	CHARACTERENDPOINT = "/lodestone/character/"
)

func (c Controller) ShallowScrapeCharacter(id int64) character.Character {
	if id == 0 {
		return character.Character{}
	}
	logger := logrus.WithField("character", id)
	logrus.Infof("Scraping Character %v", id)
	collector := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
		colly.AllowURLRevisit(),
		colly.TraceHTTP(),
	)
	var charactere character.Character
	charactere.ID = id
	logger.Info("Waiting for collector")
	collector.Wait()
	return charactere
}

func (c Controller) ScrapeCharacter(id int64, lang string) character.Character {
	if id == 0 {
		return character.Character{}
	}
	logger := logrus.WithField("character", id)
	logrus.Infof("Scraping Character %v", id)
	collector := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
		colly.AllowURLRevisit(),
		colly.TraceHTTP(),
	)
	collector.SetRequestTimeout(60 * time.Second)
	if c.proxyfunc != nil {
		logger.Info("Using Proxys for scraping characters ")
		collector.SetProxyFunc(c.proxyfunc)
	}

	// Set error handler
	collector.OnError(func(r *colly.Response, err error) {
		switch r.StatusCode {
		case 429:
			logrus.WithField("URL", r.Request.URL).Error("Too many Requests. Trying again after 2 seconds:", err)
			time.Sleep(2 * time.Second)
			err = collector.Visit(r.Request.URL.String())
			if err != nil {
				logger.Error("Visiting failed:", err)
			}
		case 0:
			logrus.WithField("URL", r.Request.URL).Error("Looks like i/o timeout. Trying again after 2 seconds:", err)
			time.Sleep(2 * time.Second)
			err = collector.Visit(r.Request.URL.String())
			if err != nil {
				logger.Error("Visiting failed:", err)
			}
		case 502:
			logrus.Error("Bad Gateway:", err)
			time.Sleep(2 * time.Second)
			err = collector.Visit(r.Request.URL.String())
			if err != nil {
				logger.Error("Visiting failed:", err)
			}
		case 404:
			logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
		case 503:
			logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, " Trying again after 2 seconds\nError:", err)
			time.Sleep(2 * time.Second)
			err = collector.Visit(r.Request.URL.String())
			if err != nil {
				logger.Error("Visiting failed:", err)
			}
		case 403:
		default:
			logrus.Error("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
		}
	})
	if c.parallel <= 0 {
		err := collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 3})
		if err != nil {
			logger.Error("Error setting Limitrule:", err)
		}
	} else {
		err := collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: c.parallel})
		if err != nil {
			logger.Error("Error setting Limitrule:", err)
		}
	}
	collector.OnRequest(func(r *colly.Request) {

		logger.Debugf("Visiting %s", r.URL.String())
		if !(strings.Contains(r.URL.String(), "friend") || strings.Contains(r.URL.String(), "achievement") || r.URL.String() == fmt.Sprintf("%v%v%d", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)) {
			r.Headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1")
		}
	})
	expire := time.Now().AddDate(0, 0, 1)
	var temp []*http.Cookie
	cookie := http.Cookie{
		Name:    "ldst_bypass_browser",
		Value:   "1",
		Expires: expire,
	}
	temp = append(temp, &cookie)
	err := collector.SetCookies(fmt.Sprintf(URL, lang), temp)
	if err != nil {
		return character.Character{}
	}
	var charactere character.Character
	charactere.ID = id
	charactere.Gearset = gear.GearSet{}
	for _, f := range characterHandlers() {
		collector.OnHTML(f(&charactere))
	}
	MAINURL := fmt.Sprintf("%v%v%d", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	CLASSURL := fmt.Sprintf("%v%v%d/class_job", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	MINIONURL := fmt.Sprintf("%v%v%d/minion", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	MOUNTURL := fmt.Sprintf("%v%v%d/mount", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	ACHIEVEMENTURL := fmt.Sprintf("%v%v%d/achievement", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	FRIENDURL := fmt.Sprintf("%v%v%d/friend", fmt.Sprintf(URL, lang), CHARACTERENDPOINT, id)
	// Set error handler
	achievements := false
	friends := false
	collector.OnHTML("li.btn__pager__current", func(e *colly.HTMLElement) {
		if (strings.Contains(e.Request.URL.String(), "achievement") && achievements) || (strings.Contains(e.Request.URL.String(), "friend") && friends) {
			return
		}
		if strings.Contains(e.Text, "Page 1 of") || strings.Contains(e.Text, "Seite 1 (von") || strings.Contains(e.Text, "Page 1 ") || strings.Contains(e.Text, "1ページ") {
			tempID, err := strconv.ParseInt(strings.ReplaceAll(After(strings.ReplaceAll(strings.ReplaceAll(e.Text, "ページ", ""), "/", " "), " "), ")", ""), 10, 0)

			if err != nil {
				logrus.Error("Error while parsing ID ", tempID)
			}
			var i int64
			var url string

			if strings.Contains(e.Request.URL.String(), "achievement") {
				url = fmt.Sprintf("%v/?page=", ACHIEVEMENTURL)
				achievements = true
			} else if strings.Contains(e.Request.URL.String(), "friend") {
				url = fmt.Sprintf("%v/?page=", FRIENDURL)
				friends = true
			}
			for i = 2; i <= tempID; i++ {
				// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				err = collector.Visit(fmt.Sprintf("%v%d", url, i))
				if err != nil {
					logger.Error("Visiting failed:", err)
				}
			}
		}
	})

	err = collector.Visit(MAINURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	err = collector.Visit(CLASSURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	err = collector.Visit(MINIONURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	err = collector.Visit(MOUNTURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	err = collector.Visit(ACHIEVEMENTURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	err = collector.Visit(FRIENDURL)
	if err != nil {
		logger.Error("Visiting failed:", err)
	}
	time.Sleep(3 * time.Second)
	logger.Info("Waiting for collector")
	collector.Wait()
	return charactere
}
