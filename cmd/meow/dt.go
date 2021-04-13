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
	if err == nil{
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nThe current progress: %s\n",meow.mName[meow.mCurrent])

		for {
			fmt.Print("> ")
			key, err := reader.ReadString( '\n' )
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			err = keyIn(key)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				if err.Error() == "bye"{
					break
				}
			}else{
				meow.mCurrent++
				if meow.mCurrent == 2{
					break
				}
			}
		}
	}
	return err
}

func keyIn(key string) error{
	key = strings.TrimRight(key, "\r\n")
	switch  {
	case meow.mKey[0] == key && meow.mCurrent == 0:
		fmt.Println(`
Meow. You have done well.

There is a box is placed
in the study. The first 
surprise will be found.
And you will have to find
the last one.

Good luck.
Regy`)
	case meow.mKey[1] == key && meow.mCurrent == 1:
		fmt.Println(`
Meow. You almost reached
the end. The next key word
is <closet>.

Thank you for your dedication
and effort. Thanks for being 
there for me as well.

Finally, don't forget to 
decrypt the file for you, meow.

Regy`)
	case key == "exit":
		return errors.New("bye")
	default:
		return errors.New("\nAuthentication failed, please verify your key.\n")
	}
	return nil
}

