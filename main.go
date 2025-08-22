package main

import (
	"fmt"

	"github.com/tmazitov/go-rsa/src"
)

func main() {

	coder := src.NewRSACoder()
	encoded := coder.Encode(3)
	fmt.Printf("Encoded: %d\n", encoded)
	decoded := coder.Decode(encoded)
	fmt.Printf("Decoded: %d\n", decoded)
}
