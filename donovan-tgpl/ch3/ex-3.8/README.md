## This doesnt work, for some reason ##

Produces black image

Anyway, compare two solutions:

### In rust ###

```rust
impl<T: Clone + Num> Add<Complex<T>> for Complex<T> {
    type Output = Complex<T>;

    #[inline]
    fn add(self, other: Complex<T>) -> Complex<T> {
	    // assuming that Complex<T>.re and Complex<T>.im implement `Add`
        Complex::new(self.re + other.re, self.im + other.im)
    }
}

// Same for Sub, Mul, ...

// So, you write:
let complex_result = (complex_a + complex_b) * complex_c;
```

### In go ###

```go
func (z *ComplexBigFloat) Add(x, y ComplexBigFloat) *ComplexBigFloat {
	tmp := big.NewFloat(0.0)
	tmp.Add(x.p, y.p)
	z.p = tmp
	tmp.Add(x.q, y.q)
	z.q = tmp

	return z
}

// Sooo.. You write:
complexRes := NewComplexBigFloat(..., ...)
complexRes.Add(complexA, complexB)
complexRes.Mul(complexRes, complexC)
```

is this looks "easier" to read?
