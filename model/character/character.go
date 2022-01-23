package character

import (
	"github.com/xivdata/goxiv/model"
	"github.com/xivdata/goxiv/model/gear"
)

type Character struct {
	ID               int64
	Name             string
	Server           *model.Server
	Achievements     []Achievement `json:",omitempty"`
	Mounts           []Mount       `json:",omitempty"`
	Minions          []Minion      `json:",omitempty"`
	Classes          []Class
	FreeCompanyID    uint64        `json:",omitempty"`
	FreeCompanyName  string        `json:",omitempty"`
	FreeCompanyURL   string        `json:",omitempty"`
	FreeCompanyCrest *model.Crest  `json:",omitempty"`
	Grandcompany     *GrandCompany `json:",omitempty"`
	Citystate        string
	CitystateID      int64 `json:",omitempty"`
	Guardian         string
	GuardianID       int64 `json:",omitempty"`
	Nameday          string
	NamedayID        int64 `json:",omitempty"`
	Tribe            string
	TribeID          int64 `json:",omitempty"`
	Sex              string
	SexID            int64 `json:",omitempty"`
	Race             string
	RaceID           int64   `json:",omitempty"`
	Bozja            *Bozja  `json:",omitempty"`
	Bio              string  `json:",omitempty"`
	Eureka           *Eureka `json:",omitempty"`
	Title            string  `json:",omitempty"`
	TitleID          int64   `json:",omitempty"`
	Friends          []*Friend
	Gearset          gear.GearSet
	PvPTeam          string      `json:",omitempty"`
	PvPTeamID        string      `json:",omitempty"`
	PvPTeamCrest     model.Crest `json:",omitempty"`
	PvPTeamURL       string      `json:",omitempty"`
	Avatar           string
	Face             string
}
