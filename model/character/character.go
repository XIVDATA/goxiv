package character

import (
	"github.com/xivdata/goxiv/model"
	"github.com/xivdata/goxiv/model/gear"
)

type Character struct {
	ID           int64
	Name         string
	Server       *model.Server
	Achievements []Achievement `json:",omitempty"`
	Mounts       []Mount       `json:",omitempty"`
	Minions      []Minion      `json:",omitempty"`
	Classes      []Class
	// FreeCompany  *freecompany.FreeCompany
	FreeCompanyID   uint64        `json:",omitempty"`
	FreeCompanyName string        `json:",omitempty"`
	FreeCompanyURL  string        `json:",omitempty"`
	Grandcompany    *GrandCompany `json:",omitempty"`
	Citystate       string
	Guardian        string
	Nameday         string
	Tribe           string
	Sex             string
	Race            string
	Bozja           *Bozja  `json:",omitempty"`
	Bio             string  `json:",omitempty"`
	Eureka          *Eureka `json:",omitempty"`
	Title           string  `json:",omitempty"`
	TitleID         int64   `json:",omitempty"`
	Friends         []*Friend
	Gearset         gear.GearSet
	PvPTeam         *PvPTeam `json:",omitempty"`
	Avatar          string
	Face            string
}
