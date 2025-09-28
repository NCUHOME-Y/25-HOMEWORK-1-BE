package main

import (
	"fmt"
)

type Complex struct {
	real      float64
	imaginary float64
}

func Add(first, second Complex) Complex {
	return Complex{
		real:      first.real + second.real,
		imaginary: first.imaginary + second.imaginary,
	}
}
func Sub(first, second Complex) Complex {
	return Complex{
		real:      first.real - second.real,
		imaginary: first.imaginary - second.imaginary,
	}
}
func Mul(first, second Complex) Complex {
	return Complex{
		real:      first.real * second.real,
		imaginary: first.imaginary * second.imaginary,
	}
}
func Div(first, second Complex) Complex {
	return Complex{
		real:      first.real / second.real,
		imaginary: first.imaginary / second.imaginary,
	}
}

// 公式参考牛顿迭代求平方根//
func ModuleLen(s Complex) float64 {
	a := s.real * s.real
	b := s.imaginary * s.imaginary
	n := a*a + b*b
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
		return fmt.Sprintf("%f", c.real)
	}
	if c.imaginary > 0 {
		return fmt.Sprintf("%f + %fi", c.real, c.imaginary)
	}
	if c.imaginary < 0 {
		return fmt.Sprintf("%f - %fi", c.real, -c.imaginary)
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

	c1 := Add(first, second)
	if c1.imaginary == 0 {
		fmt.Printf("相加为：%f\n", c1.real)
	}
	if c1.imaginary > 0 {
		fmt.Printf("相加为：%f+%fi\n", c1.real, c1.imaginary)
	}
	if c1.imaginary < 0 {
		fmt.Printf("相加为：%f-%fi\n", c1.real, -c1.imaginary)
	}

	c2 := Mul(first, second)
	if c2.imaginary == 0 {
		fmt.Printf("相乘为:%f\n", c2.real)
	}
	if c2.imaginary > 0 {
		fmt.Printf("相乘为:%f+%f\n", c2.real, c2.imaginary)
	}
	if c2.imaginary < 0 {
		fmt.Printf("相乘为:%f-%f/\n", c2.real, -c2.imaginary)
	}

	c3 := Div(first, second)
	if c3.imaginary == 0 {
		fmt.Printf("相除为:%f\n", c3.real)
	}
	if c3.imaginary > 0 {
		fmt.Printf("相除为:%f+%f\n", c3.real, c3.imaginary)
	}
	if c3.imaginary < 0 {
		fmt.Printf("相除为:%f-%f\n", c3.real, -c3.imaginary)
	}

	c4 := ModuleLen(first)
	fmt.Printf("第一个的模长：%f\n", c4)
	c5 := ModuleLen(second)
	fmt.Printf("第二个的模长：%f\n", c5)
	fmt.Println(ToString(first))
	fmt.Println(ToString(second))

}
