package main

import (
	"fmt"
	"math"
)

type Complex struct {
	real float64
	imag float64
}

func ResultComplex(real, imag float64) *Complex {
	return &Complex{real: real, imag: imag}
}

// 加法
func (num *Complex) Add(other *Complex) *Complex {
	return &Complex{
		real: num.real + other.real,
		imag: num.imag + other.imag,
	}
}

// 减法
func (num *Complex) Sub(other *Complex) *Complex {
	return &Complex{
		real: num.real - other.real,
		imag: num.imag - other.imag,
	}
}

// 乘法
func (num *Complex) Mul(other *Complex) *Complex {
	return &Complex{
		real: num.real*other.real - num.imag*other.imag,
		imag: num.real*other.imag + num.imag*other.real,
	}
}

// 除法
func (num *Complex) Div(other *Complex) *Complex {
	denominator := other.real*other.real + other.imag*other.imag
	if denominator == 0 {
		panic("division by zero")
	}

	return &Complex{
		real: (num.real*other.real + num.imag*other.imag) / denominator,
		imag: (num.imag*other.real - num.real*other.imag) / denominator,
	}
}

// 求模长
func (num *Complex) Mod() float64 {
	return math.Sqrt(num.real*num.real + num.imag*num.imag)
}

// 求共轭复数
func (num *Complex) Conjugate() *Complex {
	return &Complex{
		real: num.real,
		imag: -num.imag,
	}
}

// toString
func (num *Complex) String() string {
	if num.imag == 0 {
		return fmt.Sprintf("%g", num.real)
	}
	if num.real == 0 {
		return fmt.Sprintf("%gi", num.imag)
	}
	if num.imag > 0 {
		return fmt.Sprintf("%g+%gi", num.real, num.imag)
	}
	return fmt.Sprintf("%g%gi", num.real, num.imag)
}

func (num *Complex) GetReal() float64 {
	return num.real
}

func (num *Complex) GetImag() float64 {
	return num.imag
}

func (num *Complex) SetReal(real float64) {
	num.real = real
}

func (num *Complex) SetImag(imag float64) {
	num.imag = imag
}

// 测试
func main() {
	c1 := ResultComplex(3, 4)
	c2 := ResultComplex(1, 2)

	fmt.Printf("c1 = %s\n", c1)
	fmt.Printf("c2 = %s\n", c2)
	fmt.Printf("c1 + c2 = %s\n", c1.Add(c2))
	fmt.Printf("c1 - c2 = %s\n", c1.Sub(c2))
	fmt.Printf("c1 * c2 = %s\n", c1.Mul(c2))
	fmt.Printf("c1 / c2 = %s\n", c1.Div(c2))
	fmt.Printf("|c1| = %g\n", c1.Mod())
	fmt.Printf("c1的共轭复数 = %s\n", c1.Conjugate())
}
