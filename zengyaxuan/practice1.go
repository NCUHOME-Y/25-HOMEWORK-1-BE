package main

import (
	"errors"
	"fmt"
	"math"
)

type Complex struct {
	Real float64
	Imag float64
}

func NewComplex (real, imag float64) *Complex {
	return &Complex{Real:real, Imag:imag}
}

func (c *Complex) Add(other *Complex) *Complex {
	return &Complex{
		c.Real+other.Real,
		c.Imag+other.Imag,
	}
}

func (c *Complex) Sub(other *Complex) *Complex {
	return &Complex{
		c.Real-other.Real, 
		c.Imag-other.Imag,
	}
}

func (c *Complex) Mul(other *Complex) *Complex {
	return &Complex{
		c.Real*other.Real - c.Imag*other.Imag,
		c.Real*other.Imag + c.Imag*other.Real,
	}
}

func (c *Complex) Div(other *Complex) (*Complex,error) {
	divisor := other.Real*other.Real + other.Imag*other.Imag
	if divisor ==0 {
	return nil, errors.New("division by zero")
	}
	return &Complex{
		Real: (c.Real*other.Real + c.Imag*other.Imag)/divisor,
		Imag: (c.Imag*other.Real - c.Real*other.Imag)/divisor,
	},nil
}

func (c *Complex) Mod() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imag*c.Imag)
}

func (c *Complex)  String() string {
	if c.Imag > 0 {
		return fmt.Sprintf("%.2f+%.2fi", c.Real, c.Imag)
	}else if c.Imag == 0 {
		return fmt.Sprintf("%.2f", c.Real)
	}else{
		return fmt.Sprintf("%.2f+%.2fi", c.Real, c.Imag)
	}
}
