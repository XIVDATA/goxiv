package character

type PvPMember struct {
	Rank     string    `json:",omitempty"`
	Member   Character `json:",omitempty"`
	RankIcon string    `json:",omitempty"`
	Matches  int64     `json:",omitempty"`
}
