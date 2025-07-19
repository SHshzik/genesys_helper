package domain

type Dice map[int][]string

type Token struct {
	Count  int
	Letter string
}

const (
	Success   = "success"
	Advantage = "advantage"
)

const (
	BonusDiceLetter = "B"
)

var AvailableLetters = []string{BonusDiceLetter}

var BonusDice Dice = map[int][]string{
	1: {},
	2: {},
	3: {Success},
	4: {Success, Advantage},
	5: {Advantage, Advantage},
	6: {Advantage},
}
