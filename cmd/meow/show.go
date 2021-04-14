package meow

import "fmt"

func show()error {
	if meow.mCurrent != 2{
		fmt.Println("nothing")
	}else{
		fmt.Println("surprise!!")
	}
	return nil
}