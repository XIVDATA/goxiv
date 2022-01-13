package freecompany

type Member struct {
	Rank       string `json:",omitempty"`
	Member     int64  `json:",omitempty"`
	MemberName string `json:",omitempty"`
	MemberURL  string `json:",omitempty"`
	RankIcon   string `json:",omitempty"`
}
