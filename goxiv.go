package goxiv

import (
	"github.com/xivdata/goxiv/controller"
	"github.com/xivdata/goxiv/model/character"
	"github.com/xivdata/goxiv/model/freecompany"
	"github.com/xivdata/goxiv/model/linkshell"
	"github.com/xivdata/goxiv/model/pvpteam"
)

func ScrapeCharacter(id int64) character.Character {
	return controller.ScrapeCharacter(id)
}

func ScrapeFreecompany(id uint64) freecompany.FreeCompany {
	return controller.ScrapeFreecompany(id)
}

func ScrapePvPTeam(id string) pvpteam.PvPTeam {
	return controller.ScrapePvPTeam(id)
}

func ScrapeLinkshell(id string) linkshell.Linkshell {
	return controller.ScrapeLinkshell(id, false)
}
func ScrapeWorldLinkshell(id string) linkshell.Linkshell {
	return controller.ScrapeLinkshell(id, true)
}
