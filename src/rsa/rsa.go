package rsa

import (
	"errors"
	"math/big"
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	D *big.Int
}

type RSA struct {
	public  *PublicKey
	private *PrivateKey
	version RSAVersion
}

func NewRSACoder(version RSAVersion, e int64) (*RSA, error) {

	var E *big.Int = big.NewInt(e)

	// Negative or zero
	if E.Cmp(big.NewInt(0)) <= 0 {
		return nil, ErrRSAInvalidArg
	}

	var (
		instance *RSA = &RSA{}
		base     *rsaBase
		err      error
	)

	if base, err = setupEnvirement(version, E); err != nil {
		return nil, errors.Join(ErrRSAIntenal, err)
	}

	instance.version = version
	instance.public = &PublicKey{
		N: base.N,
		E: E,
	}
	instance.private = &PrivateKey{
		D: base.d,
	}
	return instance, nil
}
