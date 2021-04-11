package meow

import "fmt"

func help(argv []string) error{
	fmt.Println("\nUsage:  meow <command> <args>")
	fmt.Println("meow commands:")
	fmt.Println("  start     start and initialize meow playground")
	fmt.Println("  ps        list current mission progress")


	return nil
}
