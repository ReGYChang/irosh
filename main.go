package main
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"math/rand"
    "sync"

	"github.com/cheggaaa/pb"

)
func main() {
	reader := bufio.NewReader(os.Stdin)
	printMeow()
	printPb("MEOW PROGRESS",100)
	for {
		showPrompt()
		cmdString, err := reader.ReadString( '\n' )
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if cmdString == "" {
			continue
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n" )
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[ 0 ] {
		case "meow" :{
			switch arrCommandStr[ 1 ] {
				case "--help" : {
					fmt.Println("")	
					fmt.Println("Usage: meow COMMAND")	
					fmt.Println("")	
					fmt.Println("Commands:")	
					fmt.Println("  status")	
					fmt.Println("  commit")	
					fmt.Println("")
					
					return nil
				}
			}
		}
		case "exit" :
			os.Exit( 0 )
	}
	cmd := exec.Command(arrCommandStr[ 0 ], arrCommandStr[ 1 :]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
func printPb(pbName string,pbNum int){
    // create bars
    meow := pb.New(100).Prefix(pbName + ": ")
    // start pool
    pool, err := pb.StartPool(meow)
    if err != nil {
        panic(err)
    }
    // update bars
    wg := new(sync.WaitGroup)
    for _, bar := range []*pb.ProgressBar{meow} {
        wg.Add(1)
        go func(cb *pb.ProgressBar) {
            for n := 0; n < 100; n++ {
                cb.Increment()
                time.Sleep(time.Millisecond * time.Duration(rand.Intn(25)))
            }
            cb.Finish()
            wg.Done()
        }(bar)
    }
    wg.Wait()
    // close pool
    pool.Stop()
}

func showPrompt() {
    host, _ := os.Hostname()  // 取得主機名稱

    // 把字串組合起來放到 Prompt 中
    fmt.Printf("%s@%s > ", "meow", host)
}

func printMeow() {
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
