package shell

import (
	"fmt"
	"os"

)
// Print prints the prompt before waiting for user input
func promptPrint() {
	// get the host name
    host, _ := os.Hostname()  

    // combine the strings and merge into Prompt
    fmt.Printf("%s@%s > ", "meow", host)
}