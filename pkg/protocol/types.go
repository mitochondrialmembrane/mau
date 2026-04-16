package protocol

type Player interface {
	PlayCard() *Card
}

type Bot struct {
	ID   string
	Name string
	Cards []*Card
}

type Human struct {
	Cards []*Card
}

type Game struct {
	Players []*Player
}

type Card struct {
	Suit Suit
	Rank Rank
}

type Rule interface {
}

type Restriction struct {
}

type Permission struct {
}

type Suit int
type Rank int
type RuleType int

const (
	SuitDiamond   Suit = 0
	SuitClub      Suit = 1
	SuitHeart     Suit = 2
	SuitSpade     Suit = 3
	RankAce   Rank = 1
	RankTwo   Rank = 2
	RankThree Rank = 3
	RankFour  Rank = 4
	RankFive  Rank = 5
	RankSix   Rank = 6
	RankSeven Rank = 7
	RankEight Rank = 8
	RankNine  Rank = 9
	RankTen   Rank = 10
	RankJack  Rank = 11
	RankQueen Rank = 12
	RankKing  Rank = 13
)