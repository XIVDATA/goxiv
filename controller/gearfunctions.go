package controller

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/model/character"

	"github.com/xivdata/goxiv/model/gear"
)

func scrapeItem(item *gear.Item, selector string) (string, func(e *colly.HTMLElement)) {
	return selector, func(e *colly.HTMLElement) {
		tempDom, _ := e.DOM.Find("h2.db-tooltip__item__name").Html()
		if strings.Contains(tempDom, "icon") {
			item.HQ = true
		} else {
			item.HQ = true
		}
		name := e.DOM.Find("h2.db-tooltip__item__name").Children().Remove().End().Text()
		item.Name = name
		miragename := e.ChildText("div.db-tooltip__item__mirage")
		item.Mirage = miragename
		stain := e.ChildText("div.stain")
		item.Color = stain
		e.ForEach("div.db-tooltip__materia__txt", func(i int, e *colly.HTMLElement) {
			text, _ := e.DOM.Html()
			materianame := strings.Split(strings.ReplaceAll(text, "&#39;", "'"), "<br/>")[0]
			switch i {
			case 0:
				item.Materia1 = &gear.Materia{Name: materianame}
			case 1:
				item.Materia2 = &gear.Materia{Name: materianame}
			case 2:
				item.Materia3 = &gear.Materia{Name: materianame}
			case 3:
				item.Materia4 = &gear.Materia{Name: materianame}
			case 4:
				item.Materia5 = &gear.Materia{Name: materianame}
			}
		})
		e.ForEachWithBreak("div.db-tooltip__info_text", func(i int, e *colly.HTMLElement) bool {
			if strings.Contains(e.ChildAttr("a", "href"), "/lodestone/character/") {
				temp := After(BeforeLast(e.ChildAttr("a", "href"), "/"), "/")
				tempID, err := strconv.ParseInt(temp, 10, 64)
				if err != nil {
					logrus.Error("Error while parsing ID ", temp)
				}
				item.Crafter = tempID
				item.CrafterName = e.ChildText("a")
				item.CrafterURL = e.ChildAttr("a", "href")

			}
			return false
		})
	}
}

func characterGearsetClassHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	return "div.character__profile__detail", func(e *colly.HTMLElement) {
		activejob, _ := e.DOM.Find("div.character__class_icon").Children().Attr("src")
		e.ForEachWithBreak("li", func(i int, e *colly.HTMLElement) bool {
			if e.ChildAttr("img", "src") == activejob {
				name := e.ChildAttr("img", "data-tooltip")
				if name == "Blue Mage (Limited Job)" {
					name = "blue mage"
				}
				data.Gearset.Class = name
				return false
			}
			return true
		})

	}
}

func characterGearsetMainhandHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.MainHand
	item.Slot = "Mainhand"
	return scrapeItem(item, `div.icon-c--0`)
}
func characterGearsetOffhandHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.OffHand
	item.Slot = "OffHand"
	return scrapeItem(&data.Gearset.OffHand, `div.icon-c--1`)
}
func characterGearsetHeadHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Head
	item.Slot = "Head"
	return scrapeItem(&data.Gearset.Head, `div.icon-c--2`)
}
func characterGearsetBodyHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Body
	item.Slot = "Body"
	return scrapeItem(&data.Gearset.Body, `div.icon-c--3`)
}
func characterGearsetHandsHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Hands
	item.Slot = "Hands"
	return scrapeItem(&data.Gearset.Hands, `div.icon-c--4`)
}
func characterGearsetWaistHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Waist
	item.Slot = "Waist"
	return scrapeItem(&data.Gearset.Waist, `div.icon-c--5`)
}
func characterGearsetLegsHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Legs
	item.Slot = "Legs"
	return scrapeItem(&data.Gearset.Legs, `div.icon-c--6`)
}
func characterGearsetFeetHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Feet
	item.Slot = "Feet"
	return scrapeItem(&data.Gearset.Feet, `div.icon-c--7`)
}
func characterGearsetEarringHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Earring
	item.Slot = "Earring"
	return scrapeItem(&data.Gearset.Earring, `div.icon-c--8`)
}
func characterGearsetNecklaceHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Necklace
	item.Slot = "Necklace"
	return scrapeItem(&data.Gearset.Necklace, `div.icon-c--9`)
}
func characterGearsetBraceletsHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Bracelets
	item.Slot = "Bracelets"
	return scrapeItem(&data.Gearset.Bracelets, `div.icon-c--10`)
}
func characterGearsetRing1Handler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Ring1
	item.Slot = "Ring1"
	return scrapeItem(&data.Gearset.Ring1, `div.icon-c--11`)
}
func characterGearsetRing2Handler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.Ring2
	item.Slot = "Ring2"
	return scrapeItem(&data.Gearset.Ring2, `div.icon-c--12`)
}
func characterGearsetSoulCrystalHandler(data *character.Character) (string, func(e *colly.HTMLElement)) {
	item := &data.Gearset.SoulCrystal
	item.Slot = "SoulCrystal"
	return scrapeItem(&data.Gearset.SoulCrystal, `div.icon-c--13`)
}
