package meow

import (
	"fmt"
	"github.com/ReGYChang/irosh/utils"
)

func start() error {
	meow = &meowT{}

	meow.mName[0] = "MEOW  "
	meow.mName[1] = "CAESAR"
	meow.mDisc[0] = `
Meow. You have done well.

There is a box is placed
in the study. The first
surprise will be found.
And you will have to find
the last one.

Good luck.
Regy
`
	meow.mDisc[1] = `
Meow. You almost reached
the end. The next key word
is <closet>.

Thank you for your dedication
and effort. Thanks for being
there for me as well.

Finally, don't forget to
decrypt the file for you, meow.

Regy
`
	meow.mHint[0] = "this is first hint"
	meow.mHint[1] = "this is second hint"
	meow.mHint[2] = "there is no anything~"
	meow.mKey[0] = "meow"
	meow.mKey[1] = "861115"
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
