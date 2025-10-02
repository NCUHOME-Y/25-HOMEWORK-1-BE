package main
import (
	"fmt"
	"math"
)

type Complex struct {
	Real float64
	Image float64
}

func (c *Complex) Add(o Complex) Complex{
	return Complex{
		Real:c.Real + o.Real
		Image:c.Image + o.Image
	}
}

func (c *Complex) Sub(o Complex) Complex{
	return Complex{
		Real:c.Real - o.Real
		Image:c.Image - o.Image
	}
}

func (c *Complex) Mul(o Complex) Complex{
		R:= c.Real * o.Real - c.Image * o.Image
		I:= c.Real * o.Image + c.Image * o.Real 
		return Complex{Real:R , Image: I}
}

func (c *Complex) Div(o Complex) (Complex, error) {
	denominator := other.Real*other.Real + other.Imag*other.Imag
	if denominator == 0 {
		return Complex{}, fmt.Errorf("division by zero")
	}
	real := (c.Real*o.Real + c.Imag*o.Imag) / denominator
	imag := (c.Imag*o.Real - c.Real*o.Imag) / denominator
	return Complex{Real: real, Imag: imag}, nil
}

func (c *Complex) Mod() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imag*c.Imag)
}

func (c *Complex) ToString() string {
	if c.Imag >= 0 {
		return fmt.Sprintf("%.2f+%.2fi", c.Real, c.Imag)
	}
	return fmt.Sprintf("%.2f%.2fi", c.Real, c.Imag)
}
