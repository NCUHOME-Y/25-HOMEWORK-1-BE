package main

import (
	"fmt"
)

type Complex struct {
	real      float64
	imaginary float64
}

func (first Complex) Add(second Complex) Complex {
	return Complex{
		real:      first.real + second.real,
		imaginary: first.imaginary + second.imaginary,
	}
}
func (first Complex) Sub(second Complex) Complex {
	return Complex{
		real:      first.real - second.real,
		imaginary: first.imaginary - second.imaginary,
	}
}
func (first Complex) Mul(second Complex) Complex {
	return Complex{
		real:      first.real * second.real,
		imaginary: first.imaginary * second.imaginary,
	}
}
func (first Complex) Div(second Complex) Complex {
	return Complex{
		real:      first.real / second.real,
		imaginary: first.imaginary / second.imaginary,
	}
}

// 公式参考牛顿迭代求平方根//
func ModuleLen(s Complex) float64 {
	a := s.real * s.real
	b := s.imaginary * s.imaginary
	n := a + b
	if n == 0 {
		return 0
	}
	k := n
	var x float64
	for !(n-1e-10 <= x*x && x*x <= n+1e-10) {
		x = (k + n/k) / 2
		k = x
	}
	return x
}
func ToString(c Complex) string {
	if c.imaginary == 0 {
		return fmt.Sprintf("%.2f", c.real)
	}
	if c.imaginary > 0 {
		return fmt.Sprintf("%.2f + %.2fi", c.real, c.imaginary)
	}
	if c.imaginary < 0 {
		return fmt.Sprintf("%.2f - %.2fi", c.real, -c.imaginary)
	}
	return "zero"
}

func main() {
	var real1, imaginary1, real2, imaginary2 float64
	fmt.Println("请输入第一个复数的实部和虚部")
	fmt.Scan(&real1, &imaginary1)
	fmt.Println("请输入第二个复数的实部和虚部")
	fmt.Scan(&real2, &imaginary2)
	var first = Complex{real: real1, imaginary: imaginary1}
	var second = Complex{real: real2, imaginary: imaginary2}

	c1 := first.Add(second)
	fmt.Printf("相加为%s\n", ToString(c1))

	c2 := first.Sub(second)
	fmt.Printf("相减%s\n", ToString(c2))

	c3 := first.Mul(second)
	fmt.Printf("相乘为%s\n", ToString(c3))

	c4 := first.Div(second)
	fmt.Printf("相除%s\n", ToString(c4))

	c5 := ModuleLen(first)
	fmt.Printf("第一个的模长：%.2f\n", c5)
	c6 := ModuleLen(second)
	fmt.Printf("第二个的模长：%.2f\n", c6)
	fmt.Println(ToString(first))
	fmt.Println(ToString(second))
}
