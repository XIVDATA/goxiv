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

func ScrapeCharacter(id int64) character.Character {
	logrus.Infof("Scraping Character %v", id)
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
		// colly.AllowURLRevisit(),
	)
	c.SetRequestTimeout(60 * time.Second)
	// if Proxyfunc != nil {
	// 	logrus.Info("Using Proxys for scraping characters ")
	// 	c.SetProxyFunc(Proxyfunc)
	// }

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		switch r.StatusCode {
		case 429:
			logrus.WithField("id", id).WithField("URL", r.Request.URL).Error("Too many Requests. Trying again after 2 seconds:", err)
			time.Sleep(2 * time.Second)
			c.Visit(r.Request.URL.String())
		case 502:
			logrus.WithField("id", id).Error("Bad Gateway:", err)
			time.Sleep(2 * time.Second)
			c.Visit(r.Request.URL.String())
		case 404:
			logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
		case 503:
			logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
		case 403:
		default:
			logrus.Error("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
		}
	})
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})
	c.OnRequest(func(r *colly.Request) {
		if !(strings.Contains(r.URL.String(), "friend") || strings.Contains(r.URL.String(), "achievement") || r.URL.String() == fmt.Sprintf("%v%v%d", URL, CHARACTERENDPOINT, id)) {
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
	c.SetCookies(URL, temp)
	var charactere character.Character
	charactere.ID = id
	charactere.Gearset = gear.GearSet{}
	for _, f := range CharacterHandlers() {
		c.OnHTML(f(&charactere))
	}
	MAINURL := fmt.Sprintf("%v%v%d", URL, CHARACTERENDPOINT, id)
	CLASSURL := fmt.Sprintf("%v%v%d/class_job", URL, CHARACTERENDPOINT, id)
	MINIONURL := fmt.Sprintf("%v%v%d/minion", URL, CHARACTERENDPOINT, id)
	MOUNTURL := fmt.Sprintf("%v%v%d/mount", URL, CHARACTERENDPOINT, id)
	ACHIEVEMENTURL := fmt.Sprintf("%v%v%d/achievement", URL, CHARACTERENDPOINT, id)
	FRIENDURL := fmt.Sprintf("%v%v%d/friend", URL, CHARACTERENDPOINT, id)
	// Set error handler
	c.OnHTML("li.btn__pager__current", func(e *colly.HTMLElement) {
		tempID, err := strconv.ParseInt(After(e.Text, " "), 10, 0)
		if err != nil {
			logrus.Error("Error while parsing ID ", tempID)
		}
		var i int64
		var url string
		if strings.Contains(e.Request.URL.String(), "achievement") {
			url = fmt.Sprintf("%v/?page=", ACHIEVEMENTURL)
		} else if strings.Contains(e.Request.URL.String(), "friend") {
			url = fmt.Sprintf("%v/?page=", FRIENDURL)
		}
		for i = 2; i <= tempID; i++ {
			// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			c.Visit(fmt.Sprintf("%v%d", url, i))
			if err != nil {
				logrus.Error("Visiting failed:", err)
			}
		}
	})

	err := c.Visit(MAINURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	err = c.Visit(CLASSURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	err = c.Visit(MINIONURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	err = c.Visit(MOUNTURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	err = c.Visit(ACHIEVEMENTURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	err = c.Visit(FRIENDURL)
	if err != nil {
		logrus.Error("Visiting failed:", err)
	}
	c.Wait()
	logrus.Info("Waiting for collector")
	return charactere
}