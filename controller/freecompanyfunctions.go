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

func FreecompanyNameHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__name`, func(e *colly.HTMLElement) {
		data.Name = e.Text

	}

}
func FreecompanyServerHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.entry__freecompany__gc:has(i)`, func(e *colly.HTMLElement) {
		var server model.Server
		var datacenter model.Datacenter
		datacenter.Name = Between(e.Text, "(", ")")
		server.Datacenter = datacenter
		server.Name = strings.ReplaceAll(strings.ReplaceAll(BeforeLast(e.Text, "("), "\t", ""), "\n", "")

		data.Server = &server
	}

}
func FreecompanySloganHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__message`, func(e *colly.HTMLElement) {
		data.Slogan = e.Text

	}

}
func FreecompanyShortnameHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__text__tag`, func(e *colly.HTMLElement) {
		data.ShortName = strings.ReplaceAll(strings.ReplaceAll(e.Text, "«", ""), "»", "")

	}

}
func FreecompanyFoundedHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `h3.heading--lead:contains("Formed")`, func(e *colly.HTMLElement) {
		tempTime, err := strconv.ParseInt(After(BeforeLast(e.DOM.Next().Text(), ","), "("), 10, 64)
		if err != nil {
			logrus.Error("Error while parsing time for Free Company ", tempTime)
		}
		data.Founded = time.Unix(tempTime, 0)

	}

}

func FreecompanyGrandcompanyReputationHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `div.freecompany__reputation__data`, func(e *colly.HTMLElement) {
		var grandcompany freecompany.Reputation
		grandcompany.GrandCompanyName = e.ChildText("p.freecompany__reputation__gcname")
		grandcompany.Reputation = e.ChildText("p.freecompany__reputation__rank")
		data.Reputation = append(data.Reputation, grandcompany)
	}
}

func FreecompanyRankHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
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

func FreecompanyEstateHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__estate__name`, func(e *colly.HTMLElement) {
		var estate freecompany.Estate
		estate.Name = e.Text
		estate.Address = e.DOM.SiblingsFiltered("p.freecompany__estate__text").Text()
		estate.Greeting = e.DOM.SiblingsFiltered("p.freecompany__estate__greeting").Text()
		data.Estate = &estate

	}

}
func FreecompanyPvPHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("PvP")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Pvp = false
		} else {
			data.Pvp = true
		}
	}

}
func FreecompanyRaidsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Raids")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Raids = false
		} else {
			data.Raids = true
		}
	}
}
func FreecompanyTrialsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Trials")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Trials = false
		} else {
			data.Trials = true
		}
	}
}
func FreecompanyGuildhestsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Guildhests")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Guildheists = false
		} else {
			data.Guildheists = true
		}
	}
}
func FreecompanyDungeonsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Dungeons")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Dungeons = false
		} else {
			data.Dungeons = true
		}
	}
}
func FreecompanyHardcoreHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Hardcore")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Hardcore = false
		} else {
			data.Hardcore = true
		}
	}
}
func FreecompanyCasualHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Casual")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Casual = false
		} else {
			data.Casual = true
		}
	}
}
func FreecompanyLevelingHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Leveling")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Leveling = false
		} else {
			data.Leveling = true
		}
	}
}
func FreecompanyRoleplayingHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Role-playing")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Roleplay = false
		} else {
			data.Roleplay = true
		}
	}
}
func FreecompanyTankHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Tank")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Tank = false
		} else {
			data.Tank = true
		}
	}
}
func FreecompanyHealerHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Healer")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Heal = false
		} else {
			data.Heal = true
		}
	}
}
func FreecompanyDDHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("DPS")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.DD = false
		} else {
			data.DD = true
		}
	}
}
func FreecompanyGathererHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Gatherer")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Gatherer = false
		} else {
			data.Gatherer = true
		}
	}
}
func FreecompanyCrafterHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p:contains("Crafter")`, func(e *colly.HTMLElement) {
		if e.DOM.Parent().HasClass("freecompany__focus_icon--off") {
			data.Crafter = false
		} else {
			data.Crafter = true
		}
	}
}
func FreecompanyMemberHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
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
func FreecompanyAcceptsHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return `p.freecompany__recruitment`, func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "Open") {
			data.Accepts = true
		}
	}
}

func FreecompanyCrestHandler(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return "div.entry__freecompany__crest__image", func(e *colly.HTMLElement) {
		data.Crest = &freecompany.Crest{}
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

func FreecompanyHandlers() []func(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)) {
	return []func(data *freecompany.FreeCompany) (string, func(e *colly.HTMLElement)){FreecompanyCrestHandler, FreecompanyAcceptsHandler, FreecompanyGrandcompanyReputationHandler, FreecompanyMemberHandler, FreecompanyNameHandler, FreecompanyServerHandler, FreecompanySloganHandler, FreecompanyShortnameHandler, FreecompanyFoundedHandler, FreecompanyRankHandler, FreecompanyEstateHandler, FreecompanyHealerHandler, FreecompanyDDHandler, FreecompanyGathererHandler, FreecompanyCrafterHandler, FreecompanyTankHandler, FreecompanyRoleplayingHandler, FreecompanyLevelingHandler, FreecompanyPvPHandler, FreecompanyRaidsHandler, FreecompanyTrialsHandler, FreecompanyGuildhestsHandler, FreecompanyDungeonsHandler, FreecompanyHardcoreHandler, FreecompanyCasualHandler}
}
