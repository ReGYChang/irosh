package shell

import (
	"math/rand"
	"sync"
	"time"

	"github.com/cheggaaa/pb"

)

func PrintPb(pbName string, pbNum int) {
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
