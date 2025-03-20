package encode

import (
	"crypto/rand"
	"math/big"
)

func RandomeString(l int) (string, error) {
	s := make([]rune, 0, l)

	for range l {
		n, err := rand.Int(rand.Reader, big.NewInt(36))
		if err != nil {
			return "", err
		}
		r, err := IntToRune(n.Int64())
		if err != nil {
			return "", err
		}
		s = append(s, r)
	}

	return string(s), nil
}
