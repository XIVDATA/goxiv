package pvpteam

import (
	"time"

	"github.com/xivdata/goxiv/model"
)

type PvPTeam struct {
	Founded    time.Time         `json:",omitempty"`
	ID         string            `json:",omitempty"`
	Name       string            `json:",omitempty"`
	Member     []*PvPMember      `json:",omitempty"`
	Datacenter *model.Datacenter `json:",omitempty"`
	Depth      int64             `json:",omitempty"`
	MaxDepth   int64             `json:",omitempty"`
	Crest      model.Crest       `json:",omitempty"`
	URL        string            `json:",omitempty"`
}
