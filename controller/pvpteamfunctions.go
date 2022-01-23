package controller

import (
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model"
	"github.com/xivdata/goxiv/model/pvpteam"
)

func pvpTeamCrestHandler(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return "div.entry__pvpteam__crest__image", func(e *colly.HTMLElement) {
		data.Crest = model.Crest{}
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			v, _ := s.Attr("src")
			switch i {
			case 0:
				data.Crest.Top = v
			case 1:
				data.Crest.Middle = v
			case 2:
				data.Crest.Bottom = v
			}

		})
	}
}

func pvpTeamNameHandler(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return "h2.entry__pvpteam__name--team", func(e *colly.HTMLElement) {
		data.Name = e.Text
	}
}

func pvpDCHandler(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return "p.entry__pvpteam__name--dc", func(e *colly.HTMLElement) {
		data.Datacenter = &model.Datacenter{Name: e.Text}
	}
}

func pvpTeamFoundedHandler(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return `span.entry__pvpteam__data--formed`, func(e *colly.HTMLElement) {
		stringtime := After(BeforeLast(e.DOM.Find("script").Text(), ","), "(")
		tempTime, err := strconv.ParseInt(stringtime, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing time for PvP Team ", stringtime)
		}
		data.Founded = time.Unix(tempTime, 0)
	}
}

func pvpTeamMemberHandler(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return `a.entry__bg`, func(e *colly.HTMLElement) {
		var member pvpteam.PvPMember
		temp := After(BeforeLast(e.Attr("href"), "/"), "/")
		tempID, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing ID ", temp)
		}
		tempMatches, err := strconv.ParseInt(e.DOM.Find("img.entry__pvpteam__battle__icon").Siblings().Text(), 10, 64)
		if err != nil {
			logrus.Error("Error while Matches ", e.DOM.Find("img.entry__pvpteam__battle__icon").Siblings().Text())
		}
		member.Matches = tempMatches
		member.MemberID = tempID
		count := e.DOM.Find("ul.entry__freecompany__info").Children().Length()
		if count == 4 {
			member.Rank = e.DOM.Find("ul.entry__freecompany__info").Find("li").First().Text()
			tempIcon, _ := e.DOM.Find("ul.entry__freecompany__info").Find("img").First().Attr("src")
			member.RankIcon = &tempIcon
		}
		member.Member = e.ChildText("p.entry__name")
		data.Member = append(data.Member, &member)
	}
}

func pvpTeamHandlers() []func(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)) {
	return []func(data *pvpteam.PvPTeam) (string, func(e *colly.HTMLElement)){pvpTeamCrestHandler, pvpTeamMemberHandler, pvpTeamNameHandler, pvpDCHandler, pvpTeamFoundedHandler}
}
