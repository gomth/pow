package pow

import (
	"math/big"

	_ "github.com/gomth/factors"
)

func miniPow(val *big.Int, degree int) *big.Int {
	if val.Int64() == 0 {
		return big.NewInt(0)
	}

	if degree == 0 {
		return big.NewInt(1)
	}

	bigVal := val
	result := new(big.Int)
	result.Add(val, result) // copy val to result
	for i := 2; i <= degree; i++ {
		result = result.Mul(result, bigVal)
	}

	return result
}

func Get(val int, degree int) (*big.Int, error) {
	return GetUsingFactorOfPower(val, degree)
}
