package controller

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model"

	"github.com/xivdata/goxiv/model/character"
)

func characterNameHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "p.frame__chara__name", func(e *colly.HTMLElement) {
		data.Name = e.Text
	}
}
func characterTitleHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "p.frame__chara__title", func(e *colly.HTMLElement) {
		data.Title = strings.Trim(e.Text, " ")
	}
}

func characterServerDatacenterHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "p.frame__chara__world", func(e *colly.HTMLElement) {
		var server model.Server
		var datacenter model.Datacenter
		datacenter.Name = Between(e.Text, "(", ")")
		server.Datacenter = datacenter
		re := regexp.MustCompile(`[^a-zA-Z]+`)
		temp := BeforeLast(e.Text, "(")
		server.Name = re.ReplaceAllString(temp, "")
		data.Server = &server
	}

}

func characterFreecompanyHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__freecompany__name", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		temp, err := strconv.ParseUint(After(BeforeLast(link, "/"), "/"), 10, 64)
		if err == nil {
			data.FreeCompanyID = temp
			data.FreeCompanyName = e.ChildText("a")
			data.FreeCompanyURL = link
		} else {
			logrus.Error("Error while parsing Free Company: ", err, data.ID, e.Attr("href"))
		}
	}
}

func characterPvPTeamHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__pvpteam__name", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		data.PvPTeamID = After(BeforeLast(link, "/"), "/")
		data.PvPTeamURL = link
		data.PvPTeam = e.ChildText("a")
	}
}

func characterPvPTeamCrestHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__pvpteam__crest__image", func(e *colly.HTMLElement) {
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			v, _ := s.Attr("src")
			v = strings.ReplaceAll(v, "40x40", "64x64")
			data.PvPTeamCrest = model.Crest{}
			switch i {
			case 0:
				data.PvPTeamCrest.Top = v
			case 1:
				data.PvPTeamCrest.Middle = v
			case 2:
				data.PvPTeamCrest.Bottom = v
			}

		})
	}
}

func characterGrandcompanyHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `p.character-block__title:contains("Grand Company")`, func(e *colly.HTMLElement) {
		temp := e.DOM.Siblings().Text()
		var grandcompany character.GrandCompany
		grandcompany.Name = BeforeLast(temp, " /")
		grandcompany.RankName = After(temp, "/ ")
		data.Grandcompany = &grandcompany
	}
}

func characterBioHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__selfintroduction", func(e *colly.HTMLElement) {
		data.Bio = e.Text
	}
}

func characterTraitHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `p.character-block__title:contains("Race/Clan/Gender")`, func(e *colly.HTMLElement) {
		temp, _ := e.DOM.Siblings().Html()
		if strings.Contains(temp, "♀") {
			data.Sex = "♀"
		} else {
			data.Sex = "♂"
		}
		s := strings.Split(temp, "<br/>")
		data.Race = strings.ReplaceAll(s[0], "&#39;", "'")
		data.Tribe = BeforeLast(s[1], " /")

	}
}

func characterCitystageHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `p.character-block__title:contains("City-state")`, func(e *colly.HTMLElement) {
		data.Citystate = e.DOM.Siblings().Text()
	}
}

func characterNamedayHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "p.character-block__birth", func(e *colly.HTMLElement) {
		data.Nameday = e.Text
	}

}
func characterGuardianHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `p.character-block__title:contains("Guardian")`, func(e *colly.HTMLElement) {
		temp := e.DOM.SiblingsFiltered("p.character-block__name").Text()
		data.Guardian = temp
	}

}

func characterClassSpecialistHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `div.character__job__name--meister`, func(e *colly.HTMLElement) {
		level := e.DOM.SiblingsFiltered("div.character__job__level").Text()
		if level != "-" {
			exp := e.DOM.SiblingsFiltered("div.character__job__exp").Text()
			work := BeforeLast(exp, " /")
			var class character.Class
			if work == "--" && (strings.Contains(e.Text, "Blue Mage") || level == "90") {
				class.Max = true
				class.Name = e.Text
				if strings.Contains(e.Text, "Blue Mage") {
					class.Level = 70
					class.Name = "Blue Mage"
				} else {
					class.Level = 90
				}
			} else {
				work := BeforeLast(exp, " /")
				if work != "--" {
					tempexp, err := strconv.ParseInt(strings.ReplaceAll(work, ",", ""), 10, 64)
					if err != nil {
						logrus.Error("Error while parsing EXP ", work, data.ID)
					}
					class.Exp = tempexp
				} else {
					class.Exp = 0
				}
				class.Max = false
				templevel, err := strconv.ParseInt(level, 10, 64)
				if err != nil {
					logrus.Error("Error while parsing level")
				}
				class.Level = templevel
				if strings.Contains(e.Text, "Blue Mage") {
					class.Name = "Blue Mage"
				} else {
					class.Name = e.Text
				}
			}
			class.Specialist = true
			data.Classes = append(data.Classes, class)
		}
	}

}

func characterClassHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `div.character__job__name`, func(e *colly.HTMLElement) {
		level := e.DOM.SiblingsFiltered("div.character__job__level").Text()
		if level != "-" {
			exp := e.DOM.SiblingsFiltered("div.character__job__exp").Text()
			work := BeforeLast(exp, " /")
			var class character.Class
			if work == "--" && (strings.Contains(e.Text, "Blue Mage") || level == "90") {
				class.Max = true
				class.Name = e.Text
				if strings.Contains(e.Text, "Blue Mage") {
					class.Level = 70
					class.Name = "Blue Mage"
				} else {
					class.Level = 90
				}
			} else {
				work := BeforeLast(exp, " /")
				if work != "--" {
					tempexp, err := strconv.ParseInt(strings.ReplaceAll(work, ",", ""), 10, 64)
					if err != nil {
						logrus.Error("Error while parsing EXP ", work, data.ID)
					}
					class.Exp = tempexp
				} else {
					class.Exp = 0
				}
				class.Max = false
				templevel, err := strconv.ParseInt(level, 10, 64)
				if err != nil {
					logrus.Error("Error while parsing level")
				}
				class.Level = templevel
				if strings.Contains(e.Text, "Blue Mage") {
					class.Name = "Blue Mage"
				} else {
					class.Name = e.Text
				}
			}
			class.Specialist = false
			data.Classes = append(data.Classes, class)
		}
	}

}

func characterBozjaHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `div.character__job__name-sp:contains("Resistance Rank")`, func(e *colly.HTMLElement) {
		level := e.DOM.SiblingsFiltered("div.character__job__level").Text()
		exp := Between(e.DOM.SiblingsFiltered("div.character__job__exp").Text(), "Current Mettle: ", " / Mettle to Next Rank")

		templevel, err := strconv.ParseInt(strings.ReplaceAll(level, ",", ""), 10, 64)
		if err != nil {
			logrus.Error("Error while parsing Level ", level)
		}
		var temp character.Bozja
		if strings.ReplaceAll(exp, ",", "") != "--" {

			tempexp, err := strconv.ParseInt(strings.ReplaceAll(exp, ",", ""), 10, 64)
			if err != nil {
				logrus.Error("Error while parsing EXP ", exp)
			}
			temp.Mettle = tempexp
		}
		temp.Level = templevel
		data.Bozja = &temp
	}

}
func characterEurekaHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `div.character__job__name-sp:contains("Elemental Level")`, func(e *colly.HTMLElement) {
		level := e.DOM.SiblingsFiltered("div.character__job__level").Text()
		exp := BeforeLast(e.DOM.SiblingsFiltered("div.character__job__exp").Text(), " /")

		templevel, err := strconv.ParseInt(strings.ReplaceAll(level, ",", ""), 10, 64)
		if err != nil {
			logrus.Error("Error while parsing Level ", level)
		}
		var temp character.Eureka
		if strings.ReplaceAll(exp, ",", "") != "--" {

			tempexp, err := strconv.ParseInt(strings.ReplaceAll(exp, ",", ""), 10, 64)
			if err != nil {
				logrus.Error("Error while parsing EXP ", exp)
			}
			temp.Exp = tempexp
		}
		temp.Level = templevel
		data.Eureka = &temp

	}

}

func characterMinionHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `span.minion__name`, func(e *colly.HTMLElement) {
		var minion character.Minion
		minion.Name = e.Text
		data.Minions = append(data.Minions, minion)

	}

}
func characterMountHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `span.mount__name`, func(e *colly.HTMLElement) {
		var mount character.Mount
		mount.Name = e.Text
		data.Mounts = append(data.Mounts, mount)

	}

}

func characterAvatarFaceHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `img.character-block__face`, func(e *colly.HTMLElement) {
		data.Face = e.Attr("src")
	}

}

func characterAvatarHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `div.character__detail__image`, func(e *colly.HTMLElement) {
		data.Avatar = e.ChildAttr("a", "href")
	}
}

func characterAchievementHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `li.entry`, func(e *colly.HTMLElement) {
		temp := After(BeforeLast(e.ChildAttr("a.entry__achievement", "href"), "/"), "/")
		var achievement character.Achievement
		tempTime, err := strconv.ParseInt(After(BeforeLast(e.ChildText("time.entry__activity__time"), ","), "("), 10, 64)
		if err != nil {
			logrus.Error("Error while parsing time Achievement Time ", tempTime)
		}
		tempID, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing ID ", temp)
		}
		achievement.Unlocked = time.Unix(tempTime, 0)
		achievement.ID = tempID
		achievement.Name = After(BeforeLast(e.ChildText("p.entry__activity__txt"), "\""), "\"")
		data.Achievements = append(data.Achievements, achievement)
	}

}

func characterFriendHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return `a.entry__link`, func(e *colly.HTMLElement) {
		temp := After(BeforeLast(e.Attr("href"), "/"), "/")
		var friend character.Friend
		tempID, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			logrus.Error("Error while parsing ID ", temp)
		}
		friend.ID = tempID
		friend.URL = e.Attr("href")
		friend.Name = e.ChildText("p.entry__name")
		friend.Face = e.ChildAttr("img", "src")
		data.Friends = append(data.Friends, &friend)
	}

}

func characterFreecompanyCrestHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__freecompany__crest__image", func(e *colly.HTMLElement) {
		data.FreeCompanyCrest = &model.Crest{}
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			v, _ := s.Attr("src")
			v = strings.ReplaceAll(v, "40x40", "64x64")
			switch i {
			case 0:
				data.FreeCompanyCrest.Top = v
			case 1:
				data.FreeCompanyCrest.Middle = v
			case 2:
				data.FreeCompanyCrest.Bottom = v
			}

		})
	}
}

func characterHandlers() []func(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return []func(data *character.Character) (string, func(e *colly.HTMLElement)){characterPvPTeamCrestHandler, characterFreecompanyCrestHandler, characterGearsetClassHandler, characterPvPTeamHandler, characterAvatarHandler, characterGearsetOffhandHandler, characterGearsetHeadHandler, characterGearsetBodyHandler, characterGearsetHandsHandler, characterGearsetWaistHandler, characterGearsetLegsHandler, characterGearsetFeetHandler, characterGearsetEarringHandler, characterGearsetNecklaceHandler, characterGearsetRing1Handler, characterGearsetSoulCrystalHandler, characterGearsetRing2Handler, characterGearsetBraceletsHandler, characterGearsetMainhandHandler, characterAvatarFaceHandler, characterNameHandler, characterFriendHandler, characterAchievementHandler, characterTitleHandler, characterServerDatacenterHandler, characterFreecompanyHandler, characterGrandcompanyHandler, characterBioHandler, characterTraitHandler, characterCitystageHandler, characterNamedayHandler, characterGuardianHandler, characterEurekaHandler, characterBozjaHandler, characterClassHandler, characterMountHandler, characterMinionHandler, characterClassSpecialistHandler}
}
