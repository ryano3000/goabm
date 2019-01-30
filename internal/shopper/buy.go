package shopper

import (
	"math/rand"
)

// buy function is a standin for an appropriate
// buy action, it returns 1, 0 randomnly to denote a
// purchase
func buy() int {
	return rand.Intn(2)
}
