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

func (count1 Complex) Sub(count2 Complex) Complex {
	return Complex{
		Real: count1.Real - count1.Real,
		Imag: count1.Imag - count2.Imag,
	}
}
