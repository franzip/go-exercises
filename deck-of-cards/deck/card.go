package deck

type CardSign int

const (
	Spade CardSign = iota
	Heart
	Diamond
	Club
)

var Signs = []CardSign{Spade, Heart, Diamond, Club}

type Card struct {
	Value int
	Sign  CardSign
}
