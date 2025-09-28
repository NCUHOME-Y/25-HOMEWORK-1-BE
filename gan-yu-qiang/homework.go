package main

type Complex struct {
	Real float64
	Imag float64
}

func (count1 Complex) Add(count2 Complex) Complex {
	return Complex{
		Real: count1.Real + count2.Real,
		Imag: count1.Imag + count2.Imag,
	}
}

func (count1 Complex) Subtract(count2 Complex) Complex {
	return Complex{
		Real: count1.Real - count1.Real,
		Imag: count1.Imag - count2.Imag,
	}
}

func (count1 Complex) Multiply(count2 Complex) Complex {
	realPart := c1.Real*c2.Real - c1.Imag*c2.Imag
	imagPart := c1.Real*c2.Imag + c1.Imag*c2.Real
	return Complex{
		Real: realPart,
		Imag: imagPart,
	}
}
