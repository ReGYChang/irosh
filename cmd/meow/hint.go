package meow

import (
	"errors"
	"fmt"
)

func hint() error{
	if meow == nil{
		return errors.New("\nERROR: meow process have not started yet.\nPlease start the process before checking progress.\n")
	}
	fmt.Println(meow.mHint[meow.mCurrent])
	return nil
}
