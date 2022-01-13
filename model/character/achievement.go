package character

import "time"

type Achievement struct {
	ID       int64
	Unlocked time.Time
	Name     string
}
