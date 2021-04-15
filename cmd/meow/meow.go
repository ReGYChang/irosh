package meow

import (
	"errors"
)

type meowT struct {
	mProgress [3]int
	mName [3]string
	mDisc [3]string
	mHint [3]string
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
	case "hint":
		return hint()
	case "--help":
		return help(argv)
	}
}

