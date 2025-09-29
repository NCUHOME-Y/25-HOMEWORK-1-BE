package main

import (
	"fmt"
	"math"
)

type ComplexNumber interface {
	Real() float64
	Imag() float64
	Add(other ComplexNumber) ComplexNumber
	Sub(other ComplexNumber) ComplexNumber
	Mul(other ComplexNumber) ComplexNumber
	Div(other ComplexNumber) ComplexNumber
	Modulus() float64
	String() string
	
}

type Complex struct {
	real float64
	imag float64
}

func NewComplex(real, imag float64) ComplexNumber {
	return &Complex{real, imag}
}

func (c *Complex) Real() float64 {
	return c.real
}

func (c *Complex) Imag() float64 {
	return c.imag
}

func (c *Complex) Add(other ComplexNumber) ComplexNumber {
	return &Complex{
		real: c.real + other.Real(),
		imag: c.imag + other.Imag(),
	}
}

func (c *Complex) Sub(other ComplexNumber) ComplexNumber {
	return &Complex{
		real: c.real - other.Real(),
		imag: c.imag - other.Imag(),
	}
}

func (c *Complex) Mul(other ComplexNumber) ComplexNumber {
	return &Complex{
		real: c.real*other.Real() - c.imag*other.Imag(),
		imag: c.real*other.Imag() + other.Real()*c.imag,
	}
}

func (c *Complex) Div(other ComplexNumber) ComplexNumber {
	denominator := other.Real()*other.Real() + other.Imag()*other.Imag()
	if denominator == 0 {
		panic("除数不能为零")
	}
	return &Complex{
		real: (c.real*other.Real() + c.imag*other.Imag()) / denominator,
		imag: (c.imag*other.Real() - c.real*other.Imag()) / denominator,
	}
}

func (c *Complex) Modulus() float64 {
	return math.Sqrt(c.real*c.real + c.imag*c.imag)
}

func (c *Complex) String() string {
	if c.imag == 0 {
		return fmt.Sprintf("%.2f", c.real)
	}
	if c.real == 0 {
		return fmt.Sprintf("%.2fi", c.imag)
	}
	if c.imag > 0 {
		return fmt.Sprintf("%.2f+%.2fi", c.real, c.imag)
	}
	return fmt.Sprintf("%.2f%.2fi", c.real, c.imag)
}

func main() {
	var real1, imag1, real2, imag2 float64

	fmt.Print("请输入第一个复数的实部: ")
	fmt.Scan(&real1)
	fmt.Print("请输入第一个复数的虚部: ")
	fmt.Scan(&imag1)

	fmt.Print("请输入第二个复数的实部: ")
	fmt.Scan(&real2)
	fmt.Print("请输入第二个复数的虚部: ")
	fmt.Scan(&imag2)

	// 创建复数对象
	num1 := NewComplex(real1, imag1)
	num2 := NewComplex(real2, imag2)

	fmt.Printf("\n复数1: %s\n", num1)
	fmt.Printf("复数2: %s\n", num2)
	// 使用方法进行计算
	fmt.Print("加法: ")
	fmt.Println(num1.Add(num2))

	fmt.Print("减法: ")
	fmt.Println(num1.Sub(num2))

	fmt.Print("乘法: ")
	fmt.Println(num1.Mul(num2))

	fmt.Print("除法: ")
	fmt.Println(num1.Div(num2))

	fmt.Printf("复数1的模长: %.2f\n", num1.Modulus())
	fmt.Printf("复数2的模长: %.2f\n", num2.Modulus())
}
