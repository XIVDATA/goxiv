package goxiv

import (
	"github.com/xivdata/goxiv/controller"
	"github.com/xivdata/goxiv/model/character"
)

func ScrapeCharacter(id int64) character.Character {
	return controller.ScrapeCharacter(10477093)
}
