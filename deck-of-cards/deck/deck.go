package deck

type Deck []Card

const cardPerSign int = 13

func New() Deck {
	var deck []Card

	for _, sign := range Signs {
		for j := 1; j <= cardPerSign; j++ {
			deck = append(deck, Card{Value: j, Sign: sign})
		}
	}

	return deck
}

func (u Deck) sort() Deck {
	return u
}
