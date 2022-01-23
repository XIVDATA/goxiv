package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model/pvpteam"
)

const (
	PVPTEAMENDPOINT = "/lodestone/pvpteam/"
)

func (c Controller) ScrapePvPTeam(id string) pvpteam.PvPTeam {
	collector := colly.NewCollector(
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(),
	)
	collector.SetRequestTimeout(60 * time.Second)
	logrus.Infof("Scraping PvP Team %v", id)
	if c.proxyfunc != nil {
		logrus.Info("Using Proxys for scraping pvp teams")
		collector.SetProxyFunc(c.proxyfunc)
	}
	var pvpteam pvpteam.PvPTeam
	pvpteam.ID = id
	collector.OnError(func(r *colly.Response, err error) {
		collector.OnError(func(r *colly.Response, err error) {
			switch r.StatusCode {
			case 429:
				logger.WithField("URL", r.Request.URL).Error("Too many Requests. Trying again after 2 seconds:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 0:
				logger.WithField("URL", r.Request.URL).Error("Looks like i/o timeout. Trying again after 2 seconds:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 502:
				logger.Error("Bad Gateway:", err)
				time.Sleep(2 * time.Second)
				collector.Visit(r.Request.URL.String())
			case 404:
				logger.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			case 503:
				logger.Debug("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			case 403:
			default:
				logger.Error("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
			}
		})
	})
	for _, f := range pvpTeamHandlers() {
		collector.OnHTML(f(&pvpteam))
	}
	MAINURL := fmt.Sprintf("%v%v%v", URL, PVPTEAMENDPOINT, id)
	//Command to visit the website
	err := collector.Visit(MAINURL)
	if err != nil {
		log.Println("Visiting failed:", err)
	}
	logrus.Info("Waiting for Collector")
	time.Sleep(2 * time.Second)
	collector.Wait()
	return pvpteam
}
