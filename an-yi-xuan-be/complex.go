package main

import (
	"fmt"
	"math"
)

type Complex struct {
	Real float64
	Imag float64 //大写可导出
}

func NewComplex(real, imag float64) Complex {
	return Complex{Real: real, Imag: imag}
}

func (c Complex) Abs() float64 { //模长
	return math.Sqrt(c.Real*c.Real + c.Imag*c.Imag)
}

func (c Complex) Add(d Complex) Complex { //加法
	return Complex{c.Real + d.Real, c.Imag + d.Imag}
}

func (c Complex) Sub(d Complex) Complex { //减法
	return Complex{c.Real - d.Real, c.Imag - d.Imag}
}

func (c Complex) Mul(d Complex) Complex { //乘法
	return Complex{c.Real*d.Real - c.Imag*d.Imag, c.Real*d.Imag + c.Imag*d.Real}
}

func (c Complex) Div(d Complex) Complex { //除法
	denom := d.Real*d.Real + d.Imag*d.Imag
	return Complex{(c.Real*d.Real + c.Imag*d.Imag) / denom, (c.Imag*d.Real - c.Real*d.Imag) / denom} //denom为x^2+y^2
}

func (c Complex) String() string {
	const epsilon = 1e-9 // 浮点比较的容差
	switch {
	case math.Abs(c.Imag) < epsilon:
		return fmt.Sprintf("%.2f", c.Real)
	case math.Abs(c.Real) < epsilon:
		return fmt.Sprintf("%.2fi", c.Imag)
	default:
		return fmt.Sprintf("%.2f%+.2fi", c.Real, c.Imag)
	}
}

func main() {
	var r1, i1 float64
	var r2, i2 float64
	fmt.Print("请输入第一个复数的实部和虚部 (用空格隔开): ")
	_, err1 := fmt.Scanln(&r1, &i1)
	if err1 != nil {
		fmt.Println("输入格式错误:", err1)
		return // 输入错误，程序退出
	}
	fmt.Print("请输入第二个复数的实部和虚部 (用空格隔开): ")
	_, err2 := fmt.Scanln(&r2, &i2)
	if err2 != nil {
		fmt.Println("输入格式错误:", err2)
		return
	}
	c1 := NewComplex(r1, i1)
	c2 := NewComplex(r2, i2)
	fmt.Printf("你输入的两个复数是: \n c1 = %s \n c2 = %s \n", c1, c2)

	fmt.Printf("(%s) + (%s) = %s\n", c1, c2, c1.Add(c2))
	fmt.Printf("(%s) - (%s) = %s\n", c1, c2, c1.Sub(c2))
	fmt.Printf("(%s) * (%s) = %s\n", c1, c2, c1.Mul(c2))

	const epsilon = 1e-9
	if c2.Abs() < epsilon {
		fmt.Println("除数 c2 太接近于零，无法执行除法。")
	} else {
		fmt.Printf("(%s) / (%s) = %s\n", c1, c2, c1.Div(c2))
	}

	fmt.Printf("|%s| 的模长是: %.2f\n", c1, c1.Abs())
	fmt.Printf("|%s| 的模长是: %.2f\n", c2, c2.Abs())
}
