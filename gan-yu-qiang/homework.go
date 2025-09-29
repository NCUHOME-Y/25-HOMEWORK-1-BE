package main

import (
	"errors"
)

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
	realPart := count1.Real*count2.Real - count1.Imag*count2.Imag
	imagPart := count1.Real*count2.Imag + count1.Imag*count2.Real
	return Complex{
		Real: realPart,
		Imag: imagPart,
	}
}

func (count1 Complex) Divide(count2 Complex) (Complex, error) {
	mother := count2.Real*count2.Real + count2.Imag*count2.Imag
	if mother == 0 {
		var errString string = "错误：分母为零"
		return Complex{}, errors.New(errString)
	} else {
		realSon := count1.Real*count2.Real + count1.Imag*count2.Imag
		imagSon := count1.Imag*count2.Real - count1.Real*count2.Imag
		return Complex{
			Real: realSon / mother,
			Imag: imagSon / mother,
		}, nil
	}

}
