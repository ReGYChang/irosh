package cmd

import (
	"errors"

	"github.com/ReGYChang/irosh/cmd/meow"
)

// Check verifies if a command is a builtin command
func Check(argv []string) (func([]string) error, error) {
	switch argv[0] {
	case "meow":
		return meow.MeowCheck, nil
	case "exit":
		return Exit, nil
	default:
		return nil, errors.New("-meow: '" + argv[0] + "': command not found")
	}
}