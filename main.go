package main

import (
	"fmt"

	"github.com/tmazitov/go-algo/src/rsa"
)

func main() {

	var (
		base    int64  = 71
		message string = "Hello! My name is Timur!"
		client  *rsa.Client
		version rsa.RSAVersion = rsa.RSA256
		err     error
	)

	if client, err = rsa.NewClient(version, base); err != nil {
		panic(err)
	}
	fmt.Printf("Origin: %v\n", message)
	encoded, err := client.Encode([]byte(message))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded: %v\n", encoded)
	decoded, err := client.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded: %v\n", string(decoded))
}
