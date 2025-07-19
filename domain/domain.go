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
	BonusDiceLetter   = "B"
	AbilityDiceLetter = "A"
)

var AvailableLetters = []string{BonusDiceLetter, AbilityDiceLetter}

var (
	BonusDice Dice = map[int][]string{
		1: {},
		2: {},
		3: {Success},
		4: {Success, Advantage},
		5: {Advantage, Advantage},
		6: {Advantage},
	}
	AbilityDice Dice = map[int][]string{
		1: {},
		2: {Success},
		3: {Success},
		4: {Success, Success},
		5: {Advantage},
		6: {Advantage},
		7: {Success, Advantage},
		8: {Advantage, Advantage},
	}
)
