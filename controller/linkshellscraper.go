package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model/linkshell"
)

const (
	LINKSHELLENDPOINT      = "/lodestone/linkshell/"
	WORLDLINKSHELLENDPOINT = "/lodestone/crossworld_linkshell/"
)

func ScrapeLinkshell(id string, world bool) linkshell.Linkshell {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(),
	)
	c.SetRequestTimeout(60 * time.Second)
	logrus.Infof("Scraping Linkshell %v Worldlinkshell: %v", id, world)
	// if Proxyfunc != nil {
	// 	logrus.Info("Using Proxys for scraping linkshell")
	// 	c.SetProxyFunc(Proxyfunc)
	// }
	var linkshell linkshell.Linkshell
	linkshell.WorldType = world
	linkshell.ID = id
	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
	})
	for _, f := range LinkshellHandlers() {
		c.OnHTML(f(&linkshell))
	}
	var MAINURL string
	if linkshell.WorldType {
		MAINURL = fmt.Sprintf("%v%v%v", URL, WORLDLINKSHELLENDPOINT, id)
	} else {
		MAINURL = fmt.Sprintf("%v%v%v", URL, LINKSHELLENDPOINT, id)
	}

	c.OnHTML("li.btn__pager__current", func(e *colly.HTMLElement) {
		tempID, err := strconv.ParseInt(After(e.Text, " "), 10, 0)
		if err != nil {
			logrus.Error("Error while parsing ID ", tempID)
		}
		var i int64
		url := fmt.Sprintf("%v/?page=", MAINURL)

		if strings.Contains(e.Request.URL.String(), "member") {
			for i = 2; i <= tempID; i++ {
				time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
				c.Visit(fmt.Sprintf("%v%d", url, i))
				if err != nil {
					logrus.Println("Visiting failed:", err)
				}
			}
		}
	})

	err := c.Visit(MAINURL)
	if err != nil {
		logrus.Println("Visiting failed:", err)
	}
	logrus.Info("Waiting for Collector")
	c.Wait()
	return linkshell
}
