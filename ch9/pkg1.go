package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	rand "math/rand/v2"
)

func seedRand() *rand.PCG {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.NewPCG(binary.LittleEndian.Uint64(b[:]), 0)
	return r
}

func main() {
	r1 := seedRand()
	fmt.Println("r1.Uint64:", r1.Uint64())
}
