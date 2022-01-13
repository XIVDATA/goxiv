package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func ScrapeFreecompany(id uint64) freecompany.FreeCompany {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		// colly.AllowURLRevisit(),
	)
	c.SetRequestTimeout(60 * time.Second)

	// if Proxyfunc != nil {
	// 	logrus.Info("Using Proxys for scraping free companys")
	// 	c.SetProxyFunc(Proxyfunc)
	// }
	var company freecompany.FreeCompany

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL.String(), "failed with response:", r.StatusCode, "\nError:", err)
	})
	for _, f := range FreecompanyHandlers() {
		c.OnHTML(f(&company))
	}
	MAINURL := fmt.Sprintf("%v%v%d", URL, FREECOMPANYENDPOINT, id)
	MEMBERURL := fmt.Sprintf("%v%v%d/member", URL, FREECOMPANYENDPOINT, id)
	c.OnHTML("li.btn__pager__current", func(e *colly.HTMLElement) {
		tempID, err := strconv.ParseInt(After(e.Text, " "), 10, 0)
		if err != nil {
			logrus.Error("Error while parsing ID ", tempID)
		}
		var i int64
		url := fmt.Sprintf("%v/?page=", MEMBERURL)

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

	//Command to visit the website
	err := c.Visit(MAINURL)
	if err != nil {
		logrus.Println("Visiting failed:", err)
	}
	err = c.Visit(MEMBERURL)
	if err != nil {
		logrus.Println("Visiting failed:", err)
	}
	c.Wait()
	company.ID = id
	// parse our response slice into JSON format
	b, err := json.Marshal(company)
	if err != nil {
		logrus.Println("failed to serialize response:", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("fc-%d.json", company.ID), b, 0644)
	if err != nil {
		panic(err)
	}
	return company
}
