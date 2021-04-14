package meow

import (
	"errors"
	"fmt"
)

func dt(argv []string) error{
	var err error
	if len(argv) < 3{
		return errors.New("\nmissing operand after 'dt'\n")
	}else{
		switch argv[2]{
		case "meow":
			meow.mProgress[0] = 100
			meow.mCurrent++
		case "861115":
			meow.mProgress[1] = 100
			meow.mCurrent++
		default:
			err = errors.New("\nAuthentication failed, please verify your key.\n")
		}
		if meow.mCurrent == 2 {
			fmt.Println("Congrats to get all of keys! You just decrypted the top secret file, Hurry go have a look!!")
		}
	}
	err = ps(argv[:2])

	return err
}