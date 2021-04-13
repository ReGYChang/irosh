package meow

import (
	"errors"
)

type meowT struct {
	mProgress [2]int
	mName [2]string
	mDisc [2]string
	mHint [2]string
	mKey  [2]string
	mCurrent int
}

var (
	meow *meowT
)

func MeowCheck(argv []string) error {
	switch argv[1] {
	default:
		return errors.New("\nmeow: '" + argv[1] + "' is not a meow command.\nSee 'meow --help'\n")
	case "start":
		return start()
	case "ps":
		return ps(argv)
	case "dt":
		return dt(argv)
	case "show":
		return show()
	case "--help":
		return help(argv)
	}
}

