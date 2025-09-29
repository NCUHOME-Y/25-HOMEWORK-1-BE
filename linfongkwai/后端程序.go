package main

import (
	"fmt"
	"math"
)

type Complex struct {
	real float64
	imag float64
}

func Write(real, imag float64) Complex {
	return Complex{
		real: real, imag: imag,
	}
}

func (c Complex) Sum(other Complex) Complex {
	return Complex{
		real: c.real + other.real,
		imag: c.imag + other.imag,
	}
}

func (c Complex) Sub(other Complex) Complex {
	return Complex{
		real: c.real - other.real,
		imag: c.imag - other.imag,
	}
}

func (c Complex) Mul(other Complex) Complex {
	return Complex{
		real: c.real*other.real - c.imag*other.imag,
		imag: c.real*other.imag + c.imag*other.real,
	}
}

func (c Complex) Div(other Complex) Complex {
	denominator := other.real*other.real + other.imag*other.imag
	if denominator == 0 {
		return Complex{real: 0, imag: 0}
	}
	return Complex{
		real: (c.real*other.real + c.imag*other.imag) / denominator,
		imag: (c.imag*other.real - c.real*other.imag) / denominator,
	}
}

func (c Complex) Mod() float64 {
	return math.Sqrt(c.real*c.real + c.imag*c.imag)
}

func (c Complex) String() string {
	if c.imag >= 0 {
		return fmt.Sprintf("%.2f + %.2fi", c.real, c.imag)
	} else {
		return fmt.Sprintf("%.2f - %.2fi", c.real, -c.imag)
	}
}

func main() {
	var r1, r2, i1, i2 float64
	var a int

	fmt.Print("请输入第一个复数的实部:")
	fmt.Scan(&r1)
	fmt.Print("请输入第一个复数的虚部:")
	fmt.Scan(&i1)
	fmt.Print("请输入第二个复数的实部:")
	fmt.Scan(&r2)
	fmt.Print("请输入第二个复数的虚部:")
	fmt.Scan(&i2)

	c1 := Write(r1, i1)
	c2 := Write(r2, i2)

	fmt.Println("请选择要执行的步骤:")
	fmt.Println("1.加法2.减法3.乘法4.除法5.求模长6.输出字符串")
	fmt.Print("请输入:")
	fmt.Scan(&a)

	switch a {
	case 1:
		result := c1.Sum(c2)
		fmt.Printf("加法结果: %s + %s = %s", c1, c2, result)
	case 2:
		result := c1.Sub(c2)
		fmt.Printf("减法结果: %s - %s = %s", c1, c2, result)
	case 3:
		result := c1.Mul(c2)
		fmt.Printf("乘法结果: %s * %s = %s", c1, c2, result)
	case 4:
		result := c1.Div(c2)
		fmt.Printf("除法结果: %s / %s = %s", c1, c2, result)
	case 5:
		fmt.Printf("模长结果:\n")
		fmt.Printf("|%s| = %.2f\n", c1, c1.Mod())
		fmt.Printf("|%s| = %.2f", c2, c2.Mod())
	case 6:
		fmt.Printf("字符串表示:\n")
		fmt.Printf("第一个复数: %s\n", c1.String())
		fmt.Printf("第二个复数: %s", c2.String())
	default:
		fmt.Println("输入无效！")
	}
}
