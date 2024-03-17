package toolset

import (
	"crypto/rand"
	"math/big"
)

const charSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module
type Tools struct {
}

func (t *Tools) RandomString(length int) string {

	outputRunes := make([]rune, length)
	sourceRunes := []rune(charSource)

	for i := range outputRunes {
		randomIndexBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(sourceRunes))))
		if err != nil {
			continue
		}

		randomIndex := randomIndexBig.Int64()
		outputRunes[i] = sourceRunes[randomIndex]
	}

	return string(outputRunes)
}
