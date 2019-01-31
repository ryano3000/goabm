package shopper

import (
	"math/rand"
	"time"
)

// buy function is a standin for an appropriate
// buy action, it returns 1, 0 randomnly to denote a
// purchase
func buy() int {
	time.Sleep(time.Second)
	return rand.Intn(2)
}
