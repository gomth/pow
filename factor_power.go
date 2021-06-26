package pow

import (
	"math/big"

	"github.com/gomth/factors"
)

// GetUsingFactorOfPower - treat degree as multiplication of prime numbers
func GetUsingFactorOfPower(val int, degree int) (*big.Int, error) {
	factors, _ := factors.Get(degree)

	res := big.NewInt(int64(val))
	for _, degreeMul := range factors {
		res = miniPow(res, degreeMul)
	}

	return res, nil
}
