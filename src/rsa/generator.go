package rsa

import (
	"math/big"
)

// isGeneratedValid checks is generated values are valid for RSA algorithm.
func isGeneratedValid(p, q, e *big.Int) bool {

	var (
		one   *big.Int = big.NewInt(1)
		mone  *big.Int = big.NewInt(-1)
		gcd   *big.Int = big.NewInt(0)
		euler *big.Int = big.NewInt(0)
	)

	if p.Cmp(q) == 0 {
		return false
	}

	if !p.ProbablyPrime(20) || !q.ProbablyPrime(20) {
		return false
	}

	// Values aren't prime to each other
	if gcd.GCD(nil, nil, p, q).Cmp(one) != 0 {
		return false
	}

	// Values (p-1)(q-1) and e aren't prime to each other
	euler.Mul(p.Add(p, mone), q.Add(q, mone))
	gcd.GCD(nil, nil, euler, e)
	if gcd.Cmp(one) != 0 {
		return false
	}
	p.Add(p, one)
	q.Add(q, one)
	return true
}

type rsaBase struct {
	d *big.Int // private key (for decoding)
	N *big.Int // public key (for encoding)
}

// setupEnvirement prepare all values for the future encoding/decoding.
func setupEnvirement(version RSAVersion, e *big.Int) (*rsaBase, error) {

	var (
		p    *big.Int
		q    *big.Int
		base *rsaBase = &rsaBase{
			d: nil,
			N: nil,
		}
		extra      *big.Int = big.NewInt(0)
		mone       *big.Int = big.NewInt(-1)
		randomizer *Randomizer
		err        error
	)

	// 1. Using custom randomizer, generate 2 random prime values
	// p and q. They shouldn't be equal and should be relatively prime to 'e' value. 
	if randomizer, err = NewRandomizer(int(version)); err != nil {
		return nil, err
	}
	defer randomizer.Close()

	for {
		if p, err = randomizer.GenerateBigInt(); err != nil {
			return nil, err
		}

		if q, err = randomizer.GenerateBigInt(); err != nil {
			return nil, err
		}

		if (isGeneratedValid(p, q, e)) {
			break
		}
	}

	// 2. Calculate N value as a public key for the future encoding.
	base.N = big.NewInt(0).Mul(p, q)

	// 3. Calculate d value as a private key for the future decoding
	euler := big.NewInt(0).Mul(p.Add(p, mone), q.Add(q, mone))
	base.d = big.NewInt(0)
	big.NewInt(0).GCD(base.d, extra, e, euler)
	base.d = base.d.Mod(base.d, euler)

	return base, nil
}
