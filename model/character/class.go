package character

type Class struct {
	ID    int64 `json:",omitempty"`
	Level int64
	Exp   int64
	Max   bool
	Name  string
}
