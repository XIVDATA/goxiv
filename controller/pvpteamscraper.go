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

func ScrapePvPTeam(id string) pvpteam.PvPTeam {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(),
	)
	c.SetRequestTimeout(60 * time.Second)
	logrus.Infof("Scraping PvP Team %v", id)
	// if Proxyfunc != nil {
	// 	log.Info("Using Proxys for scraping pvp teams")
	// 	c.SetProxyFunc(Proxyfunc)
	// }
	var pvpteam pvpteam.PvPTeam
	pvpteam.ID = id
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
	})
	for _, f := range PvPTeamHandlers() {
		c.OnHTML(f(&pvpteam))
	}
	MAINURL := fmt.Sprintf("%v%v%v", URL, PVPTEAMENDPOINT, id)
	//Command to visit the website
	err := c.Visit(MAINURL)
	if err != nil {
		log.Println("Visiting failed:", err)
	}
	logrus.Info("Waiting for Collector")
	c.Wait()
	return pvpteam
}
