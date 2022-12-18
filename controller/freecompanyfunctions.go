package controller

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model"
	"github.com/xivdata/goxiv/model/freecompany"
)

const (
	LEADERICON = "https://img.finalfantasyxiv.com/lds/h/Z/W5a6yeRyN2eYiaV-AGU7mJKEhs.png"
)

func freecompanyNameHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__name`, func(e *colly.HTMLElement) {
		data.Name = e.Text

	}

}
func freecompanyServerHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.entry__freecompany__gc:has(i)`, func(e *colly.HTMLElement) {
		var server model.Server
		var datacenter model.Datacenter
		datacenter.Name = Between(e.Text, "(", ")")
		server.Datacenter = datacenter
		server.Name = strings.ReplaceAll(strings.ReplaceAll(BeforeLast(e.Text, "("), "\t", ""), "\n", "")

		data.Server = &server
	}

}
func freecompanySloganHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__message`, func(e *colly.HTMLElement) {
		data.Slogan = e.Text

	}

}
func freecompanyShortnameHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__tag`, func(e *colly.HTMLElement) {
		data.ShortName = strings.ReplaceAll(strings.ReplaceAll(e.Text, "«", ""), "»", "")

	}

}
func freecompanyFoundedHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `h3.heading--lead:contains("Formed")`, func(e *colly.HTMLElement) {
		tempTime, err := strconv.ParseInt(After(BeforeLast(e.DOM.Next().Text(), ","), "("), 10, 64)
		if err != nil {
			logrus.Error("Error while parsing time for Free Company ", tempTime)
		}
		data.Founded = time.Unix(tempTime, 0)

	}

}

func freecompanyGrandcompanyReputationHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `div.freecompany__reputation__data`, func(e *colly.HTMLElement) {
		var grandcompany freecompany.Reputation
		grandcompany.GrandCompanyName = e.ChildText("p.freecompany__reputation__gcname")
		grandcompany.Reputation = e.ChildText("p.freecompany__reputation__rank")
		data.Reputation = append(data.Reputation, grandcompany)
	}
}

func freecompanyRankHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `h3.heading--lead:contains("Rank")`, func(e *colly.HTMLElement) {
		if e.Text == "Rank" {
			rank, err := strconv.ParseInt(e.DOM.NextFiltered("p.freecompany__text").Text(), 10, 64)
			if err != nil {
				logrus.Error("Error while parsing rank ")
			}
			data.Rank = rank
		}

	}

}

func freecompanyEstateHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__estate__name`, func(e *colly.HTMLElement) {
		var estate freecompany.Estate
		estate.Name = e.Text
		estate.Address = e.DOM.SiblingsFiltered("p.freecompany__estate__text").Text()
		estate.Greeting = e.DOM.SiblingsFiltered("p.freecompany__estate__greeting").Text()
		data.Estate = &estate

	}

}
func freecompanyPvPHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("PvP")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Pvp = false
		} else {
			data.Pvp = true
		}
	}

}
func freecompanyRaidsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Raids")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Raids = false
		} else {
			data.Raids = true
		}
	}
}
func freecompanyTrialsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Trials")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Trials = false
		} else {
			data.Trials = true
		}
	}
}
func freecompanyGuildhestsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Guildhests")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Guildheists = false
		} else {
			data.Guildheists = true
		}
	}
}
func freecompanyDungeonsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Dungeons")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Dungeons = false
		} else {
			data.Dungeons = true
		}
	}
}
func freecompanyHardcoreHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Hardcore")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Hardcore = false
		} else {
			data.Hardcore = true
		}
	}
}
func freecompanyCasualHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Casual")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Casual = false
		} else {
			data.Casual = true
		}
	}
}
func freecompanyLevelingHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Leveling")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Leveling = false
		} else {
			data.Leveling = true
		}
	}
}
func freecompanyRoleplayingHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Role-playing")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Roleplay = false
		} else {
			data.Roleplay = true
		}
	}
}
func freecompanyTankHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Tank")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Tank = false
		} else {
			data.Tank = true
		}
	}
}
func freecompanyHealerHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Healer")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Heal = false
		} else {
			data.Heal = true
		}
	}
}
func freecompanyDDHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("DPS")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.DD = false
		} else {
			data.DD = true
		}
	}
}
func freecompanyGathererHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Gatherer")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Gatherer = false
		} else {
			data.Gatherer = true
		}
	}
}
func freecompanyCrafterHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Crafter")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Crafter = false
		} else {
			data.Crafter = true
		}
	}
}
func freecompanyMemberHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `a.entry__bg`, func(e *colly.HTMLElement) {
		temp := After(BeforeLast(e.Attr("href"), "/"), "/")
		var fcmember freecompany.Member
		tempID, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing ID ", temp)
		}
		fcmember.Member = tempID
		fcmember.MemberName = e.ChildText("p.entry__name")
		fcmember.MemberURL = e.Attr("href")
		fcmember.Rank = e.DOM.Find("ul.entry__freecompany__info").Find("li").First().Text()
		fcmember.RankIcon, _ = e.DOM.Find("ul.entry__freecompany__info").Find("img").First().Attr("src")
		data.Members = append(data.Members, &fcmember)
		if fcmember.RankIcon == LEADERICON {
			data.Leader = tempID
			data.LeaderName = e.ChildText("p.entry__name")
			data.LeaderURL = e.Attr("href")
		}
	}

}
func freecompanyAcceptsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__recruitment`, func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Open") {
			data.Accepts = true
		}
	}
}

func freecompanyCrestHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return "div.entry__freecompany__crest__image", func(e *colly.HTMLElement) {
		data.Crest = &model.Crest{}
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			v, _ := s.Attr("src")
			switch i {
			case 2:
				data.Crest.Top = v
			case 1:
				data.Crest.Middle = v
			case 0:
				data.Crest.Bottom = v
			}

		})
	}
}

func freecompanyHandlers() []func(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return []func(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)){freecompanyCrestHandler, freecompanyAcceptsHandler, freecompanyGrandcompanyReputationHandler, freecompanyMemberHandler, freecompanyNameHandler, freecompanyServerHandler, freecompanySloganHandler, freecompanyShortnameHandler, freecompanyFoundedHandler, freecompanyRankHandler, freecompanyEstateHandler, freecompanyHealerHandler, freecompanyDDHandler, freecompanyGathererHandler, freecompanyCrafterHandler, freecompanyTankHandler, freecompanyRoleplayingHandler, freecompanyLevelingHandler, freecompanyPvPHandler, freecompanyRaidsHandler, freecompanyTrialsHandler, freecompanyGuildhestsHandler, freecompanyDungeonsHandler, freecompanyHardcoreHandler, freecompanyCasualHandler}
}
