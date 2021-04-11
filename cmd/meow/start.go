package meow

import (
	"fmt"
	"github.com/ReGYChang/irosh/utils"
)

func start() error {
	meow = &meowT{}

	meow.mDisc[0] = "this is first question"
	meow.mDisc[1] = "this is second question"
	meow.mHint[0] = "this is first hint"
	meow.mHint[1] = "this is second hint"

	utils.PrintPb("Initializing process: ",100,100)
	fmt.Print("meow: process initialization" + " succeeded[92mGreen[0m")

	return nil
}
