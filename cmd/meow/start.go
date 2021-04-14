package meow

import (
	"errors"
	"fmt"
	"github.com/ReGYChang/irosh/utils"
)

func start() error {
	if meow != nil{
		return errors.New("\nmeow has already been initiated")
	}else{
		meow = &meowT{}

		meow.mName[0] = "MISSION 1"
		meow.mName[1] = "MISSION 2"
		meow.mName[2] = "FINISHED"
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
		meow.mDisc[2] = "FINISHED"
		meow.mHint[0] = "Just press meow into prompt (do not forget to meow dt <ARGS>)"
		meow.mHint[1] = "keeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		meow.mHint[2] = "FINISHED"
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
	}

	return nil
}
