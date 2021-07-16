package gacha

type Rarity string

const (
	RarityN  Rarity = "N"
	RarityR  Rarity = "R"
	RaritySR Rarity = "SR"
	RarityXR Rarity = "XR"
)

func (r Rarity) String() string {
	return string(r)
}

type Card struct {
	// データストアでrarityという名前になるようにタグを設定する
	Rarity Rarity `datastore:"rarity"`
	Name   string `datastore:"name"`
}

func (c *Card) String() string {
	return c.Rarity.String() + ":" + c.Name
}
