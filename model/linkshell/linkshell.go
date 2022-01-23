package linkshell

import (
	"time"

	"github.com/xivdata/goxiv/model"
)

type Linkshell struct {
	ID         string
	Members    []LinkshellMember
	Name       string
	Datacenter *model.Datacenter `json:",omitempty"`
	Founded    time.Time
	WorldType  bool
}
