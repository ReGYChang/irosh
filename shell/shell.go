package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ReGYChang/irosh/cmd/meow"
)

var (
	reader *bufio.Reader
)

// Run contains the entire shell lifecycle
func Run() {
	// initialize the shell
	initialize()

	// interpret each input indefinetly
	for {
		err := interpret()
		if err == nil {
			continue
		}
		if err == io.EOF {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		fmt.Println(err)
	}

}

func initialize() {
	reader = bufio.NewReader(os.Stdin)
	meow.Init()

	fmt.Println("               ....                          .... 	")
	fmt.Println("              ..x....                      ....x.. 	")
	fmt.Println("             ..xx......     ........     ......xx.. 	")
	fmt.Println("            ..xxxx...,,. .............. .,,...xxxx.. 	")
	fmt.Println("            ..xxxxx,,,,..................,,,,xxxxx.. 	")
	fmt.Println("             .,,,,..,,...................,,..,,,,,. 	")
	fmt.Println("           ........ ,,,.................,,, ......... 	")
	fmt.Println("         ....... .(((,,,...............,,,))). ........ 	")
	fmt.Println("        ..... ..,,a@@@@a,,...........,,a@@@@a,,.. ...... 	")
	fmt.Println("       .......,,a@@`  '@@,...........,@@`  '@@a,,........ 	")
	fmt.Println("       .......,,@@@    @@@,.a@@@@@a.,@@@    @@@,,........ 	")
	fmt.Println("       ....,,,,,,@@@aa@@@,,,,`@@@',,,,@@@aa@@@,,,,,,,.... 	")
	fmt.Println("        ...,,,,,,,,,,,,,,,,,,,,|,,,,,,,,,,,,,,,,,,,,,... 	")
	fmt.Println("          ...,,,,,,,,,,,,,,,,`   ',,,,,,,,,,,,,,,,,... 	")
	fmt.Println("              .. ,,,,,,,,,,,,,...,,,,,,,,,,,,,, .. 	")
	fmt.Println("      (     ......... ,,,,,,,,,,,,,,,,,,, ........... 		")
	fmt.Println("   (   )  .............._ _ _ _ _ _ _ _................   ( 		")
	fmt.Println("    )  ( ...............................................   ) 	")
	fmt.Println("   (   ) ...............................................  (  ) 	")
	fmt.Println("    ) ( ,,,,,,,,,,,,,,, ................. ,,,,,,,,,,,,,,,, ) ( 		")
	fmt.Println(" ,%%%%,,,,,,,,,,,,,,,,,, ............... ,,,,,,,,,,,,,,,,,,%%%%, 	")
	fmt.Println(" %%%%%`.,,(,,(,,(,,(,,'%%%%%%%%%%%%%%%%%%`,,,),,),,),,),,.'%%%%% 	")
	fmt.Println(" `%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%' 	")
	fmt.Println("    %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% 	")
	fmt.Println("    ::::::;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::::: 	")
	fmt.Println("   ::::::;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::::: 	")
	fmt.Println("  ::::::;;%%;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::::: 	")
	fmt.Println(" ::::::;;%%;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;%%:::::: 	")
	fmt.Println("::::::;;%%%;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;%%::::::")
	fmt.Println("::::::;;%%%;;;;A;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;%%%:::::")
	fmt.Println("::::::;;;%%;;;;;AA;;;;;;;;;;;;;;;;;;;;A;;;;;;;;;;;;;;;;;;%%%:::::")
	fmt.Println("::::::;;;;%%;;;;;AAA;;;;;;;;;;;;;;;;AA;;;;;;;;;;;;A;;;;;;%%::::::")
	fmt.Println("::::::;;A;;;;;;;;;AAA;;;;;;;;;A;;;;AAA;;;;;;;;;;;;;AA;;;%%;::::::")
	fmt.Println(" ::::::;AA;;;;;;;;;AAA;;;;;;;A;;;;;AAAA;;;;;A;;;;;;AAA;;;;:::::: 	")
	fmt.Println("  ::::::;AAA;;;;;;;AAA;;A;;;AA;;;;;;AAAA;;;;AA;;;;;AAA;;;:::::: 	")
	fmt.Println("    :::::;AAA;;;;;AAA;;AA;;;AAA;;;;;;AAAA;;AAA;;;;AAAA;;::::: 	")
	fmt.Println("       :::;AAAA;;AAAA;;AAA;;;AAA;;;;AAAAA;AAA;;;;AAAAAA::: 	")
	fmt.Println("          ::AAAAAAAA;;;;AAA;AAAAA;;AAAAA;;;AAA;;AAAAAAA 	")
	fmt.Println("            .::::::                           ::::::. 	")
	fmt.Println("           :::::::'                           `:::::::	")
	fmt.Println("/***************************************************************/")
	fmt.Println("/* This is good luck meow meow ~                               */")
	fmt.Println("/* Welcome to my playground ~~~                                */")
	fmt.Println("/* Run 'meow --help' for more information on a command.        */")
	fmt.Println("/***************************************************************/")
	fmt.Println("")
	fmt.Println("")
}
