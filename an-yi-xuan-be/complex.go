package main

import (
	"fmt"
	"math"
)

type Complex struct {
	shi float64
	xu  float64
}

func NewComplex(shi, xu float64) Complex {
	return Complex{shi: shi, xu: xu}
}

func (c Complex) Abs() float64 { //模长
	return math.Sqrt(c.shi*c.shi + c.xu*c.xu)
}

func (c Complex) Add(d Complex) Complex { //加法
	return Complex{c.shi + d.shi, c.xu + d.xu}
}

func (c Complex) Sub(d Complex) Complex { //减法
	return Complex{c.shi - d.shi, c.xu - d.xu}
}

func (c Complex) Mul(d Complex) Complex { //乘法
	return Complex{c.shi*d.shi - c.xu*d.xu, c.shi*d.xu + c.xu*d.shi}
}

func (c Complex) Div(d Complex) Complex { //除法
	denom := d.shi*d.shi + d.xu*d.xu
	return Complex{(c.shi*d.shi + c.xu*d.xu) / denom, (c.xu*d.shi - c.shi*d.xu) / denom} //denom为x^2+y^2
}

func (c Complex) String() string {
	const epsilon = 1e-9 // 浮点比较的容差
	switch {
	case math.Abs(c.xu) < epsilon:
		return fmt.Sprintf("%.2f", c.shi)
	case math.Abs(c.shi) < epsilon:
		return fmt.Sprintf("%.2fi", c.xu)
	default:
		return fmt.Sprintf("%.2f%+.2fi", c.shi, c.xu)
	}
}

func main() {
	var s1, x1 float64
	var s2, x2 float64
	fmt.Print("请输入第一个复数的实部和虚部 (用空格隔开): ")
	_, err1 := fmt.Scanln(&s1, &x1)
	if err1 != nil {
		fmt.Println("输入格式错误:", err1)
		return // 输入错误，程序退出
	}
	fmt.Print("请输入第二个复数的实部和虚部 (用空格隔开): ")
	_, err2 := fmt.Scanln(&s2, &x2)
	if err2 != nil {
		fmt.Println("输入格式错误:", err2)
		return
	}
	c1 := NewComplex(s1, x1)
	c2 := NewComplex(s2, x2)
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
