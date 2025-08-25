package rsa

import (
	"math/big"
	"os"
)

type Randomizer struct {
	reader *os.File
	length int
}

// Randomizer is a structer, that can generate random number with fixed length (in bytes) using system file /dev/urandom.
func NewRandomizer(length int) (*Randomizer, error) {

	var (
		instance *Randomizer = &Randomizer{length: length}
		err      error
	)

	if instance.reader, err = os.Open("/dev/urandom"); err != nil {
		return nil, err
	}

	return instance, nil
}

func (r *Randomizer) Close() {
	r.reader.Close()
}

// GenerateBigInt generates a big.Int value with fixed length (in bytes).
func (r *Randomizer) GenerateBigInt() (*big.Int, error) {

	var (
		err         error
		buffer      []byte   = make([]byte, r.length/8)
		randomValue *big.Int = big.NewInt(0)
	)

	if _, err = r.reader.Read(buffer); err != nil {
		return nil, err	
	}

	randomValue.SetBytes(buffer)
	randomValue.Abs(randomValue)
	return randomValue, nil
}
