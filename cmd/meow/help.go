package meow

import "fmt"

func help(argv []string) error{
	fmt.Println(`
Usage:  meow COMMAND |<ARGS>|

meow commands:
  start     start and initialize meow playground
  ps        list current mission progress

 `)
	fmt.Println("\\033[32m大家好\\033[0m")

	return nil
}
