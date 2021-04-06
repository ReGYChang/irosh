package shell

import (
	"fmt"
	"os"

)

// Prompt contains what needs to be displayed before asking for user input
// type Prompt struct {
// 	Items []Item
// }

// Item represents how a prompt item should behave
// type Item interface {
// 	Generate() Item
// 	Prefix()
// 	String() string
// }

// Print prints the prompt before waiting for user input
func promptPrint() {
	// 取得主機名稱
    host, _ := os.Hostname()  

    // 把字串組合起來放到 Prompt 中
    fmt.Printf("%s@%s > ", "meow", host)
}

// Generate initializes all prompt items
// func (p *Prompt) Generate() {
// 	for i := 0; i < len(p.Items); i++ {
// 		p.Items[i] = p.Items[i].Generate()
// 	}
// }
