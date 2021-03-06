package bufbig_test

import (
	"github.com/farces/dumb/bufbig"
	//"./bufbig"
	"math/big"
	"testing"
)

//addition test
func TestBigAccumulatorAdd(t *testing.T) {
	out := big.NewInt(int64(10))

	v := bufbig.NewBigAccumulator()
	v.AddInt(2)
	v.AddInt(8)

	if v.Value().String() != out.String() {
		t.Errorf("Addition: 2 + 8 = %v, want %v", v.Value().String(), out.String())
	}
}

//subtraction test
func TestBigAccumulatorSub(t *testing.T) {
	out := big.NewInt(int64(-10))

	v := bufbig.NewBigAccumulator()
	v.AddInt(2)
	v.AddInt(-12)

	if v.Value().String() != out.String() {
		t.Errorf("Subtraction: 2 - 12 = %v, want %v", v.Value().String(), out.String())
	}
}

//setstring test
func TestBigAccumulatorSetString(t *testing.T) {
	out := big.NewInt(int64(12345654321))

	v := bufbig.NewBigAccumulator()
	v.SetString("12345654321", 10)

	if v.Value().String() != out.String() {
		t.Errorf("SetValue(\"12345654321\",10) = %v, want %v", v.Value().String(), out.String())
	}
}

//setstring invalid value test
func TestBigAccumulatorSetStringInvalid(t *testing.T) {
	out := big.NewInt(int64(1))

	v := bufbig.NewBigAccumulator()
	v.AddInt(1)
	res := v.SetString("boobs", 10)

	//expecting SetValue to pass SetString success result back
	if res != false {
		t.Errorf("SetValue Invalid = true, want false")
	}

	//checking to make sure original value persists
	if v.Value().String() != out.String() {
		t.Errorf("SetValue(\"boobs\",10) = %v, want %v", v.Value().String(), out.String())
	}
}

//setbigint test
func TestBigAccumulatorSetBigInt(t *testing.T) {
	out := big.NewInt(int64(12345))

	v := bufbig.NewBigAccumulator()
	v.SetBigInt(big.NewInt(int64(12345)))

	if v.Value().String() != out.String() {
		t.Errorf("SetBigInt big.NewInt(int64(12345)) = %v, want %v", v.Value().String(), out.String())
	}
}

//reset test
func TestBigAccumulatorReset(t *testing.T) {
	out := bufbig.NewBigAccumulator()

	v := bufbig.NewBigAccumulator()
	v.AddInt(10)
	v.Reset()

	if v.Value().String() != out.Value().String() {
		t.Errorf("Reset() = %v, want %v", v.Value().String(), out.Value().String())
	}
}

//int64 overflow test
func TestBigAccumulatorOverflow(t *testing.T) {
	out, _ := new(big.Int).SetString("9223372036854775808", 10)

	v := bufbig.NewBigAccumulator()
	v.SetString("9223372036854775807", 10)
	v.AddInt(1)

	if v.Value().String() != out.String() {
		t.Errorf("MaxInt64 + 1 = %v, want %v", v.Value().String(), out.String())
	}
}

//int64 underflow test
func TestBigAccumulatorUnderflow(t *testing.T) {
	out, _ := new(big.Int).SetString("-9223372036854775808", 10)

	v := bufbig.NewBigAccumulator()
	v.SetString("-9223372036854775807", 10)
	v.AddInt(-1)

	if v.Value().String() != out.String() {
		t.Errorf("MinInt64 - 1 = %v, want %v", v.Value().String(), out.String())
	}
}
