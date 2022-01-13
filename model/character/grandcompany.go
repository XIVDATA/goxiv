package character

type GrandCompany struct {
	ID       int64 `json:",omitempty"`
	Rank     int64 `json:",omitempty"`
	Name     string
	RankName string
}
