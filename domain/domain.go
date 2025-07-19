package domain

type Dice map[int][]string

type Token struct {
	Count  int
	Letter string
}

const (
	Success   = "success"
	Advantage = "advantage"
	Triumph   = "triumph"

	Failure      = "failure"
	Complication = "complication"
	Crash        = "crash"
)

const (
	BonusDiceLetter   = "B"
	AbilityDiceLetter = "A"
	MasterDiceLetter  = "M"

	PenaltyDiceLetter    = "P"
	DifficultyDiceLetter = "D"
	ChallengeDiceLetter  = "C"
)

var AvailableLetters = []string{BonusDiceLetter, AbilityDiceLetter, MasterDiceLetter, PenaltyDiceLetter, DifficultyDiceLetter, ChallengeDiceLetter}

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
	MasterDice Dice = map[int][]string{
		1:  {},
		2:  {Success},
		3:  {Success},
		4:  {Success, Success},
		5:  {Success, Success},
		6:  {Advantage},
		7:  {Success, Advantage},
		8:  {Success, Advantage},
		9:  {Success, Success},
		10: {Advantage, Advantage},
		11: {Advantage, Advantage},
		12: {Triumph},
	}

	PenaltyDice Dice = map[int][]string{
		1: {},
		2: {},
		3: {Failure},
		4: {Failure},
		5: {Complication},
		6: {Complication},
	}
	DifficultyDice Dice = map[int][]string{
		1: {},
		2: {Failure},
		3: {Failure, Failure},
		4: {Complication},
		5: {Complication},
		6: {Complication},
		7: {Complication, Complication},
		8: {Failure, Complication},
	}
	ChallengeDice Dice = map[int][]string{
		1:  {},
		2:  {Failure},
		3:  {Failure},
		4:  {Failure, Failure},
		5:  {Failure, Failure},
		6:  {Complication},
		7:  {Complication},
		8:  {Failure, Complication},
		9:  {Failure, Complication},
		10: {Complication, Complication},
		11: {Complication, Complication},
		12: {Crash},
	}
)
