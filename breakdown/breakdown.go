package breakdown

import (
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

type Breakdown struct {
	Value    float32
	Sign     int
	Exponent int
	Mantissa uint32
	// MantissaSummation is how to compute how the mantissa is used in the calculation.
	// The mantissa is expressed as an integer, but the computation is done via a summation that
	// can be seen in the for loop of the breakdown. This gets the value to be <0.
	MantissaSummation decimal.Decimal
}

func (b Breakdown) String() string {
	bits := math.Float32bits(b.Value)
	sign := bits >> (23 + 8)
	exponent := (bits >> 23) & 0b11111111
	mantissa := bits & 0b11111111111111111111111

	var s strings.Builder
	s.WriteString(fmt.Sprintf("%10s: %10.5f 0b%032b\n", "Binary", b.Value, bits))
	s.WriteString(fmt.Sprintf("%10s: %11d b%01b\n", "Sign", sign, sign))
	s.WriteString(fmt.Sprintf("%10s: %11d 0b%01b\n", "Exponent", int(exponent)-127, exponent))
	s.WriteString(fmt.Sprintf("%10s: %11d 0b%01b\n", "Mantissa", mantissa, mantissa))
	s.WriteString("\n")
	s.WriteString("Throw this into google to get your number.\n")
	s.WriteString(fmt.Sprintf("(-1)^%d * %s * 2^(%d)\n", sign, b.MantissaSummation.String(), int(exponent)-127))
	return s.String()
}

// BreakdownFloat breaks down the components of the IEEE 754 float32.
// https://en.wikipedia.org/wiki/IEEE_754
func BreakdownFloat(f float32) Breakdown {
	bits := math.Float32bits(f)
	// 1 Sign
	// 8 Exponent
	// 23 Mantissa

	sign := bits >> (23 + 8)
	exponent := (bits >> 23) & 0b11111111
	mantissa := bits & 0b11111111111111111111111

	next := mantissa
	sum := decimal.NewFromInt(1)
	// This takes each bit, and take it's value * 2^-i.
	// Usually binary is expressed as 2^i position, but we want to get the value as <0,
	// so we flip this on it's head. This can be done with normal integers, but I was lazy and just used a decimal
	// library. It also helps me print the decimal number as a precise decimal.
	for i := 1; i <= 23; i++ {
		v := (next & 0b10000000000000000000000) >> 22
		d := decimal.NewFromInt(int64(v))
		mag := decimal.NewFromInt(2).Pow(decimal.NewFromInt(-1 * int64(i)))
		b := mag.Mul(d)
		sum = sum.Add(b)
		next = next << 1
	}

	return Breakdown{
		Value:             f,
		Sign:              int(sign),
		Exponent:          int(exponent) - 127,
		Mantissa:          mantissa,
		MantissaSummation: sum,
	}
}
