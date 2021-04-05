package cmd

import(
	"errors"

	"github.com/ReGYChang/irosh/meow"
)

// Check verifies if a command is a builtin command
func Check(argv []string) (func([]string) error, error) {
	switch argv[0] {
	case "exit":
		return Exit(), nil
	default:
		return nil, errors.New(argv[0] + " is not a meow command. See 'meow --help'.")
	}
}