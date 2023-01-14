package decks

type Deck struct {
	Name      string
	Commander NameCountPair
	Companion NameCountPair
	Sideboard []NameCountPair
	Deck      []NameCountPair
}

type NameCountPair struct {
	Count int
	Name  string
}
