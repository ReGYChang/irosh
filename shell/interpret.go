package shell

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ReGYChang/irosh/cmd"

)

func interpret() error {

	// print the prompt
	promptPrint()

	// read user input
	input, err := readInput()

	// return errors
	if err != nil {
		if err == io.EOF {
			return err
		}
		return err
	}

	input = strings.TrimRight(input, "\r\n")

	// if no input is given, skip the cycle
	if input == "" {
		return nil
	}

	// separate the input in arguments
	argv := strings.Fields(input)

	// check if the command is a builtin command
	fn, err := cmd.Check(argv)
	if err == nil {
		err = fn(argv)
		return err
	}

	// otherwise, execute the command
	command := exec.Command("bash", "-c", input)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()

}

func readInput() (string, error) {
	return reader.ReadString('\n')
}
