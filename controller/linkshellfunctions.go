package controller

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model"
	"github.com/xivdata/goxiv/model/linkshell"
)

func LinkshellMemberHandler(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)) {
	return "a.entry__link", func(e *colly.HTMLElement) {
		var member linkshell.LinkshellMember
		tempString := After(BeforeLast(e.Attr("href"), "/"), "/")
		tempID, err := strconv.ParseInt(tempString, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing ID ", tempString)
		}
		member.CharacterID = tempID
		temp := e.DOM.Find("div.entry__chara_info__linkshell")
		rankicon, exists := temp.Find("img").Attr("src")
		if exists {
			member.Rank = temp.Text()
			member.RankIcon = &rankicon
		}
		member.Character = e.ChildText("p.entry__name")
		data.Members = append(data.Members, member)

	}
}

func LinkshellNameHandler(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)) {
	return "h3.heading__linkshell__name", func(e *colly.HTMLElement) {
		data.Name = e.DOM.Children().Remove().End().Text()
	}
}

func WorldLinkshellDatacenterHandler(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)) {
	return "span.heading__cwls__dcname", func(e *colly.HTMLElement) {
		if data.WorldType {
			data.Datacenter = &model.Datacenter{Name: e.Text}
		}
	}
}
func WorldLinkshellFoundedHandler(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)) {
	return "span.heading__cwls__formed", func(e *colly.HTMLElement) {
		stringtime := After(BeforeLast(e.DOM.Find("script").Text(), ","), "(")
		tempTime, err := strconv.ParseInt(stringtime, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing time for Linkshell ", stringtime)
		}
		data.Founded = time.Unix(tempTime, 0)
	}
}

func LinkshellHandlers() []func(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)) {
	return []func(data *linkshell.Linkshell) (string, func(e *colly.HTMLElement)){WorldLinkshellFoundedHandler, WorldLinkshellDatacenterHandler, LinkshellNameHandler, LinkshellMemberHandler}
}
