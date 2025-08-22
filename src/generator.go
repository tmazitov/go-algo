package src

import (
	"math/big"
	"math/rand"
)

func isGeneratedValid(p, q, e *big.Int) bool {

	var (
		one   *big.Int = big.NewInt(1)
		mone  *big.Int = big.NewInt(-1)
		gcd   *big.Int = big.NewInt(0)
		euler *big.Int = big.NewInt(0)
	)

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

func setupEnvirement() (*big.Int, *big.Int, *big.Int, *big.Int) {

	var (
		p     *big.Int = big.NewInt(13)
		q     *big.Int = big.NewInt(4)
		e     *big.Int = big.NewInt(5)
		d     *big.Int = big.NewInt(0)
		extra *big.Int = big.NewInt(0)
		one   *big.Int = big.NewInt(1)
		mone  *big.Int = big.NewInt(-1)
	)
	if p.Sign() == -1 {
		p.Abs(p)
	}
	if q.Sign() == -1 {
		q.Abs(q)
	}

	for p.Cmp(q) == 0 || !isGeneratedValid(p, q, e) {
		q = big.NewInt(rand.Int63())
		p = big.NewInt(rand.Int63())
		if p.Sign() == -1 {
			p.Abs(p)
		}
		if q.Sign() == -1 {
			q.Abs(q)
		}
	}

	euler := big.NewInt(0)
	euler.Mul(p.Add(p, mone), q.Add(q, mone))

	p.Add(p, one)
	q.Add(q, one)

	gcd := big.NewInt(0)
	gcd.GCD(d, extra, e, euler)
	d = d.Mod(d, euler)

	return p, q, e, d
}
