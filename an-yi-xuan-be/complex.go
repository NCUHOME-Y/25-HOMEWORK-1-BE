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
	return Complex{(c.shi*d.shi + c.xu*d.xu) / denom, (c.xu*d.shi - c.shi*d.xu) / denom}
}

func (c Complex) String() string {
	// 根据虚部的正负来决定使用 + 还是 -     且为0实部只输出0
	if c.xu == 0 {
		return fmt.Sprintf("%.2f", c.shi)
	} else if c.xu > 0 {
		return fmt.Sprintf("%.2f + %.2fi", c.shi, c.xu)
	}
	return fmt.Sprintf("%.2f - %.2fi", c.shi, -c.xu)
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
	fmt.Println("\n--------------------")
	fmt.Printf("你输入的两个复数是: \n c1 = %s \n c2 = %s \n", c1, c2)
	fmt.Println("--------------------")

	// --- 执行并打印计算结果 ---
	fmt.Printf("(%s) + (%s) = %s\n", c1, c2, c1.Add(c2))
	fmt.Printf("(%s) - (%s) = %s\n", c1, c2, c1.Sub(c2))
	fmt.Printf("(%s) * (%s) = %s\n", c1, c2, c1.Mul(c2))

	// 检查除数是否为0
	if c2.shi == 0 && c2.xu == 0 {
		fmt.Println("除数 c2 是零，无法执行除法。") //float64不可以直接判断
	} else {
		fmt.Printf("(%s) / (%s) = %s\n", c1, c2, c1.Div(c2))
	}

	fmt.Printf("|%s| 的模长是: %.2f\n", c1, c1.Abs())
	fmt.Printf("|%s| 的模长是: %.2f\n", c2, c2.Abs())
}
