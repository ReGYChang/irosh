package meow

import (
	"errors"
	"github.com/ReGYChang/irosh/utils"
)

func ps(argv []string) error {
	if meow == nil{
		return errors.New("\nERROR: meow process have not started yet.\nPlease start the process before checking progress.\n")
	}else if len(argv) == 2 {
		utils.PrintPb(meow.mName[0],meow.mProgress[0],25)
		utils.PrintPb(meow.mName[1],meow.mProgress[1],25)
	}else{
		return errors.New("\nmeow: meow ps accepts no arguments.\nSee 'meow --help.'\n")
	}
	return nil
}