package meow

import "fmt"

func help(argv []string) error{
	fmt.Println(`
Usage:  meow COMMAND |<ARGS>|

meow commands:
  start     start and initialize meow playground
  ps        list current mission progress
  dt        key in all the keys to decrypt the top secret file
  hint      you will get some hints about the keys
  show      show the top secret file

 `)
	return nil
}
