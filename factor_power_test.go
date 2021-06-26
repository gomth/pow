package pow

import (
	"math/big"
	"testing"
)

func TestFDPowNegEval(t *testing.T) {
	val, err := GetUsingFactorOfPower(-5, 24)
	if err != nil {
		t.Errorf("TestminiPowZero failed with unexpected error.")
	}

	if val.Cmp(big.NewInt(59604644775390625)) != 0 {
		t.Errorf("TestminiPowZero failed.\nExpected: %v\nActual: %v", "59604644775390625", val)
	}
}

func TestFDPowNegOdd(t *testing.T) {
	val, err := GetUsingFactorOfPower(-5, 27)
	if err != nil {
		t.Errorf("TestminiPowZero failed with unexpected error.")
	}

	expected := big.Int{}
	expected.SetString("-7450580596923828125", 10)

	if expected.Cmp(val) != 0 {
		t.Errorf("TestminiPowNegOdd failed.\nExpected: %v\nActual: %v", "-7450580596923828125", val)
	}
}

func TestDFPowPositiveBig(t *testing.T) {
	val, err := GetUsingFactorOfPower(3, 110)
	if err != nil {
		t.Errorf("TestminiPowZero failed with unexpected error.")
	}

	expected := big.Int{}
	expected.SetString("30432527221704537086371993251530170531786747066637049", 10)

	if expected.Cmp(val) != 0 {
		t.Errorf("TestminiPowPositiveBig failed.\nExpected: %v\nActual: %v", "30432527221704537086371993251530170531786747066637049", val)
	}
}
