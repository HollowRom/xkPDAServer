package netTools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandNumber() uint64 {
	return rand.Uint64()
}
