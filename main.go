package main

import (
	"fmt"

	"github.com/tmazitov/go-algo/src/rsa"
)

func main() {

	var (
		base    int64  = 71
		message string = "Hello world!"
		coder   *rsa.RSA
		version rsa.RSAVersion = rsa.RSA128
		err     error
	)

	if coder, err = rsa.NewRSACoder(version, base); err != nil {
		panic(err)
	}
	fmt.Printf("Origin: %s\n", message)
	encoded := coder.Encode(message)
	fmt.Printf("Encoded: %d\n", encoded)
	decoded := coder.Decode(encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
