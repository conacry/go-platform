package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	defaultMinNumber = 0
	defaultMaxNumber = 100
)

func RandomDefaultNumber() int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(defaultMaxNumber-defaultMinNumber))
	if err != nil {
		panic(err)
	}

	return int(defaultMinNumber + nBig.Int64())
}

func RandomNumber(minNumber, maxNumber int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxNumber-minNumber))
	if err != nil {
		panic(err)
	}

	return int(minNumber + nBig.Int64())
}
