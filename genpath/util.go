package genpath

import "math/rand"

func randInt(from, to int) int {
	return int(rand.Int31n(int32(to)) + int32(from))
}
