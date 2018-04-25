package frac

import (
	"math/big"
)

func Abs(z ComplexBigFloat) big.Float {
	// math.Sqrt(p*p + q*q)
	res := big.NewFloat(0.0)
	res.Mul(z.p, z.q)
	res.Sqrt(res)

	return *res
}

type ComplexBigFloat struct {
	p *big.Float
	q *big.Float
}

func NewComplexBigFloat() ComplexBigFloat {
	return ComplexBigFloat{
		big.NewFloat(0.0),
		big.NewFloat(0.0),
	}
}

func (z *ComplexBigFloat) Mul(x, y ComplexBigFloat) *ComplexBigFloat {
	// real = (x.p * y.p - x.q * y.q)
	// imag = (x.p * y.q + x.q * y.p)

	tmp1 := big.NewFloat(0.0)
	tmp2 := big.NewFloat(0.0)

	tmp1.Mul(x.p, y.p)
	tmp2.Mul(x.q, y.q)
	z.p.Sub(tmp1, tmp2)

	tmp1.Mul(x.p, y.q)
	tmp2.Mul(x.q, y.p)
	z.q.Add(tmp1, tmp2)

	return z
}

func (z *ComplexBigFloat) Add(x, y ComplexBigFloat) *ComplexBigFloat {
	tmp1 := big.NewFloat(0.0)
	tmp1.Add(x.p, y.p)
	z.p = tmp1
	tmp1.Add(x.q, y.q)
	z.q = tmp1

	return z
}
