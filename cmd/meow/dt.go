package meow

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func dt(argv []string) error{
	err := ps(argv)
	if err != nil{
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nThe current progress is <%s>\n",meow.mName[meow.mCurrent])

		for {
			fmt.Print("> ")
			key, err := reader.ReadString( '\n' )
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			err = keyIn(key)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}else{
				meow.mCurrent++
				if meow.mCurrent == 3{
					fmt.Println("Congrats to get all of keys!")
					os.Exit(0)
				}
			}
		}
	}else{
		return err
	}
	return nil
}

func keyIn(key string) error{
	key = strings.TrimSuffix(key, "\n" )
	if key != meow.mKey[meow.mCurrent]{
		return errors.New("\nkey incorrect!\n")
	}
	return nil
}

