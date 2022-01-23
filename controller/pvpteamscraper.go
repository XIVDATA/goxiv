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
		fmt.Println("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
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
	collector.Wait()
	return pvpteam
}
