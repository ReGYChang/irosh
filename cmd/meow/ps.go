package meow

import(
	"github.com/ReGYChang/irosh/utils"
	"errors"
)

func Ps(argv []string) error {
	if len(argv) == 2 {
		utils.PrintPb("Mission 1",meow.mProgress[0])
		utils.PrintPb("Mission 2",meow.mProgress[1])
	}else{
		return errors.New("meow: meow ps accepts no arguments.\nSee 'meow --help.'")
	}
	return nil
}