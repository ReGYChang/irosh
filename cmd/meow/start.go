package meow

import (
	"fmt"
	"github.com/ReGYChang/irosh/utils"
)

func start() error {
	meow = &meowT{}

	meow.mName[0] = "  MEOW"
	meow.mName[1] = "CAESAR"
	meow.mDisc[0] = "this is first question"
	meow.mDisc[1] = "this is second question"
	meow.mHint[0] = "this is first hint"
	meow.mHint[1] = "this is second hint"
	meow.mKey[0] = "meow"
	meow.mKey[1] = "caesar"
	meow.mCurrent = 0

	utils.PrintPb("Initializing process",100,100)
	fmt.Println("meow:process initialization succeeded!")
	fmt.Println(`

-----BEGIN PGP SIGNED MESSAGE-----
Hash:SHA1

Hello Garra. Happy birthday to you.

To celebrate, I have devised a test.

There is a top secret file encrypted
with AES256. You should pass the test
and get the key. It will lead you on 
the road to solve the puzzle.

Good luck, meow.

Regy

-----END PGP SIGNED MESSAGE-----
 `)

	return nil
}
