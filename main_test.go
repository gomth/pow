package pow

import (
	"math/big"
	"testing"
)

func TestMiniPowZero(t *testing.T) {
	val := miniPow(big.NewInt(0), 1024)

	if val.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("TestminiPowZero failed.\nExpected: %v\nActual: %v", 0, val)
	}
}

func TestMiniPowZeroDegree(t *testing.T) {
	val := miniPow(big.NewInt(1024), 0)

	if val.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("TestminiPowZero failed.\nExpected: %v\nActual: %v", 1, val)
	}
}

func TestMiniPowNegEval(t *testing.T) {
	val := miniPow(big.NewInt(-5), 2)

	if val.Cmp(big.NewInt(25)) != 0 {
		t.Errorf("TestminiPowZero failed.\nExpected: %v\nActual: %v", 25, val)
	}
}

func TestMiniPowNegOdd(t *testing.T) {
	val := miniPow(big.NewInt(-5), 5)

	expected := big.Int{}
	expected.SetString("-3125", 10)

	if expected.Cmp(val) != 0 {
		t.Errorf("TestminiPowNegOdd failed.\nExpected: %v\nActual: %v", "-3125", val)
	}
}

func TestMiniPowPositiveBig(t *testing.T) {
	val := miniPow(big.NewInt(3), 7)

	expected := big.Int{}
	expected.SetString("2187", 10)

	if expected.Cmp(val) != 0 {
		t.Errorf("TestminiPowPositiveBig failed.\nExpected: %v\nActual: %v", "2187", val)
	}
}

func TestMiniPowGet(t *testing.T) {
	val, err := Get(3, 7)
	if err != nil {
		t.Errorf("TestMiniPowGet failed.\nUnexpected error: %v\n", err)
	}

	expected := big.Int{}
	expected.SetString("2187", 10)

	if expected.Cmp(val) != 0 {
		t.Errorf("TestminiPowPositiveBig failed.\nExpected: %v\nActual: %v", "2187", val)
	}
}
