package freecompany

type Reputation struct {
	GrandCompany     int64  `json:",omitempty"`
	Reputation       string `json:",omitempty"`
	GrandCompanyName string `json:",omitempty"`
}
