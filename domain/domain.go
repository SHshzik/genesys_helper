package domain

type Dice map[int][]string

const (
	Success   = "success"
	Advantage = "advantage"
)

var BonusDice Dice = map[int][]string{
	1: {},
	2: {},
	3: {Success},
	4: {Success, Advantage},
	5: {Advantage, Advantage},
	6: {Advantage},
}
