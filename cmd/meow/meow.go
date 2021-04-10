package meow

import (
	"errors"
)

type meowT struct {
	mProgress [2]int
	mDisc [2]string
	mHint [2]string
}

var (
	meow meowT
)

func MeowCheck(argv []string) error {
	switch argv[1] {
		default:
			return errors.New("meow: " + argv[1] + " is not a meow command.\nSee 'meow --help'")
		case "ps":
			return Ps(argv)
	}
}

func Init(){
	meow = meowT{}

	meow.mDisc[0] = "this is first question"
	meow.mDisc[1] = "this is second question"
	meow.mHint[0] = "this is first hint"
	meow.mHint[1] = "this is second hint"
}