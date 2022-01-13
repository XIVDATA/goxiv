package gear

type GearSet struct {
	Class       string
	ClassID     int64
	MainHand    Item `json:",omitempty"`
	OffHand     Item `json:",omitempty"`
	Head        Item `json:",omitempty"`
	Body        Item `json:",omitempty"`
	Hands       Item `json:",omitempty"`
	Waist       Item `json:",omitempty"`
	Legs        Item `json:",omitempty"`
	Feet        Item `json:",omitempty"`
	Earring     Item `json:",omitempty"`
	Necklace    Item `json:",omitempty"`
	Bracelets   Item `json:",omitempty"`
	Ring1       Item `json:",omitempty"`
	Ring2       Item `json:",omitempty"`
	SoulCrystal Item `json:",omitempty"`
}
