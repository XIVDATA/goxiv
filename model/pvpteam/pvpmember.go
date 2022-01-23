package pvpteam

type PvPMember struct {
	Rank     string `json:",omitempty"`
	Member   string `json:",omitempty"`
	MemberID int64  `json:",omitempty"`
	RankIcon string `json:",omitempty"`
	Matches  int64  `json:",omitempty"`
}
