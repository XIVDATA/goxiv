package freecompany

import (
	"time"

	"github.com/xivdata/goxiv/model"
)

type FreeCompany struct {
	ID          uint64
	Name        string
	Founded     time.Time
	Server      *model.Server
	Crest       *model.Crest
	ShortName   string
	Slogan      string
	Rank        int64
	Accepts     bool
	Leader      int64
	LeaderName  string
	LeaderURL   string
	Members     []*Member
	Estate      *Estate `json:",omitempty"`
	Roleplay    bool
	Leveling    bool
	Casual      bool
	Hardcore    bool
	Dungeons    bool
	Guildheists bool
	Trials      bool
	Raids       bool
	Pvp         bool
	Tank        bool
	Heal        bool
	DD          bool
	Crafter     bool
	Gatherer    bool
	Reputation  []Reputation
}
