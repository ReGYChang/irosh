package meow

import "fmt"

func hint() error{
	fmt.Println(meow.mHint[meow.mCurrent])
	return nil
}
