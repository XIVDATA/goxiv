package goxiv

import (
	"github.com/xivdata/goxiv/controller"
	"github.com/xivdata/goxiv/model/character"
	"github.com/xivdata/goxiv/model/freecompany"
)

func ScrapeCharacter(id int64) character.Character {
	return controller.ScrapeCharacter(id)
}

func ScrapeFreecompany(id uint64) freecompany.FreeCompany {
	return controller.ScrapeFreecompany(id)
}
