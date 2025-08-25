package rsa

import "math/big"

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	D *big.Int
}

type KeyPair struct {
	Private *PrivateKey
	Public  *PublicKey
}
