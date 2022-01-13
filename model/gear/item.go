package gear

type Item struct {
	ID          int64 `json:",omitempty"`
	Name        string
	Mirage      string   `json:",omitempty"`
	MirageID    int64    `json:",omitempty"`
	Crafter     int64    `json:",omitempty"`
	CrafterName string   `json:",omitempty"`
	CrafterURL  string   `json:",omitempty"`
	Materia1    *Materia `json:",omitempty"`
	Materia2    *Materia `json:",omitempty"`
	Materia3    *Materia `json:",omitempty"`
	Materia4    *Materia `json:",omitempty"`
	Materia5    *Materia `json:",omitempty"`
	Color       string   `json:",omitempty"`
	ColorID     int64    `json:",omitempty"`
	Slot        string
	HQ          bool
}
