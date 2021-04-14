package meow

import (
	"errors"
	"fmt"
)

func show()error {
	if meow == nil{
		return errors.New("\nERROR: meow process have not started yet.\nPlease start the process before checking progress.\n")
	}else if meow.mCurrent != 2{
		fmt.Println("nothing")
	}else{
		fmt.Println("surprise!!")
	}
	return nil
}