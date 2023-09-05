package generator

import (
	"crypto/rand"
	"math/big"
)

func RandomNumber(minNumber, maxNumber int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxNumber-minNumber))
	if err != nil {
		panic(err)
	}

	return int(minNumber + nBig.Int64())
}
