package toolset

import (
	"crypto/rand"
	"math/big"
)

// Tools is the type used to instantiate this module
type Tools struct {
}

func (t *Tools) RandomPassword(length int) string {
	var outputRunes = make([]rune, length)
	charRangeStart, charRangeEnd := 33, 126

	for i := 0; i < length; i++ {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(charRangeEnd-charRangeStart+1)))
		outputRunes[i] = rune(int64(charRangeStart) + charIndex.Int64())
	}

	return string(outputRunes)
}
