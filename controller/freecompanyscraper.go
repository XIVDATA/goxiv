package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model/freecompany"
)

const (
	URL                 = "http://eu.finalfantasyxiv.com"
	FREECOMPANYENDPOINT = "/lodestone/freecompany/"
)

func (c Controller) ScrapeFreecompany(id uint64) freecompany.FreeCompany {
	collector := colly.NewCollector(
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(),
	)
	collector.SetRequestTimeout(60 * time.Second)
	logrus.Infof("Scraping Free Company %v", id)
	if c.proxyfunc != nil {
		logrus.Info("Using Proxys for scraping free companys")
		collector.SetProxyFunc(c.proxyfunc)
	}
	var company freecompany.FreeCompany

	// Set error handler
	collector.OnError(func(r *colly.Response, err error) {

		collector.OnError(func(r *colly.Response, err error) {
			switch r.StatusCode {
			case 429:
				logrus.WithField("URL", r.Request.URL).Error("Too many Requests. Trying again after 2 seconds:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 0:
				logrus.WithField("URL", r.Request.URL).Error("Looks like i/o timeout. Trying again after 2 seconds:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 502:
				logrus.Error("Bad Gateway:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 404:
				logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			case 503:
				logrus.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			case 403:
			default:
				logrus.Error("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			}
		})
	})
	for _, f := range freecompanyHandlers() {
		collector.OnHTML(f(&company))
	}
	MAINURL := fmt.Sprintf("%v%v%d", URL, FREECOMPANYENDPOINT, id)
	MEMBERURL := fmt.Sprintf("%v%v%d/member", URL, FREECOMPANYENDPOINT, id)
	collector.OnHTML("li.btn__pager__current", func(e *colly.HTMLElement) {
		tempID, err := strconv.ParseInt(After(e.Text, " "), 10, 0)
		if err != nil {
			logrus.Error("Error while parsing ID ", tempID)
		}
		var i int64
		url := fmt.Sprintf("%v/?page=", MEMBERURL)

		if strings.Contains(e.Request.URL.String(), "member") {
			for i = 2; i <= tempID; i++ {
				time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
				collector.Visit(fmt.Sprintf("%v%d", url, i))
				if err != nil {
					logrus.Println("Visiting failed:", err)
				}
			}
		}
	})

	//Command to visit the website
	err := collector.Visit(MAINURL)
	if err != nil {
		logrus.Println("Visiting failed:", err)
	}
	err = collector.Visit(MEMBERURL)
	if err != nil {
		logrus.Println("Visiting failed:", err)
	}
	logrus.Info("Waiting for Collector")
	time.Sleep(2 * time.Second)
	collector.Wait()

	company.ID = id
	return company
}
