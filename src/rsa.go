package src

import "math/big"

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	D *big.Int
}

type RSA struct {
	p       *big.Int
	q       *big.Int
	public  *PublicKey
	private *PrivateKey
}

func NewRSACoder() *RSA {
	var (
		instance   *RSA     = &RSA{}
		p, q, e, d *big.Int = setupEnvirement()
		N          *big.Int = p.Mul(p, q)
	)

	instance.p = p
	instance.q = q
	instance.public = &PublicKey{
		N: N,
		E: e,
	}
	instance.private = &PrivateKey{
		D: d,
	}
	return instance
}
