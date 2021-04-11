package meow

import (
	"errors"
	"github.com/ReGYChang/irosh/utils"
)

func ps(argv []string) error {
	if meow == nil{
		return errors.New("meow: meow process have not started yet.\nStart the process before checking progress.")
	}else if len(argv) == 2 {
		utils.PrintPb("Mission 1",meow.mProgress[0],25)
		utils.PrintPb("Mission 2",meow.mProgress[1],25)
	}else{
		return errors.New("meow: meow ps accepts no arguments.\nSee 'meow --help.'")
	}
	return nil
}