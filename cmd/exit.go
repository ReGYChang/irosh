package cmd

import (
	"errors"
	"os"

)

func Exit() error{
	os.Exit(0)

	return nil
}