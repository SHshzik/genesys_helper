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
)

const (
	BonusDiceLetter   = "B"
	AbilityDiceLetter = "A"
	MasterDiceLetter  = "M"
)

var AvailableLetters = []string{BonusDiceLetter, AbilityDiceLetter, MasterDiceLetter}

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
)
