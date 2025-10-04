package main

import (
	"fmt"
	"math"
)

// 面向对象
type Complex struct {
	real float64
	imag float64
}

func NewComplex(real, imag float64) Complex {
	return Complex{real: real, imag: imag}
}

func (a Complex) Add(b Complex) Complex {
	return NewComplex(a.real+b.real, a.imag+b.imag)
}

func (a Complex) Subtract(b Complex) Complex {
	return NewComplex(a.real-b.real, a.imag-b.imag)
}

func (a Complex) Multiply(b Complex) Complex {
	real := a.real*b.real - a.imag*b.imag
	imag := a.real*b.imag + a.imag*b.real
	return NewComplex(real, imag)
}

func (a Complex) Divide(b Complex) Complex {
	under := b.real*b.real + b.imag*b.imag
	real := (a.real*b.real + a.imag*b.imag) / under
	imag := (a.imag*b.real - a.real*b.imag) / under
	return NewComplex(real, imag)
}

func (a Complex) Magnitude() float64 {
	return math.Sqrt(a.real*a.real + a.imag*a.imag)
}

func (a Complex) String() string {
	return fmt.Sprintf("(%.2f + %.2fi)", a.real, a.imag)
}

func main() {
	var choice int
	var r1, i1, r2, i2 float64

	fmt.Println("=== 面向对象复数计算器 ===")

	fmt.Print("请输入第一个复数的实部: ")
	fmt.Scan(&r1)
	fmt.Print("请输入第一个复数的虚部: ")
	fmt.Scan(&i1)

	fmt.Print("请输入第二个复数的实部: ")
	fmt.Scan(&r2)
	fmt.Print("请输入第二个复数的虚部: ")
	fmt.Scan(&i2)

	c1 := NewComplex(r1, i1)
	c2 := NewComplex(r2, i2)

	fmt.Printf("\n第一个复数: %s\n", c1)
	fmt.Printf("第二个复数: %s\n", c2)

	fmt.Println("\n请选择运算:")
	fmt.Println("1. 加法 (+)")
	fmt.Println("2. 减法 (-)")
	fmt.Println("3. 乘法 (×)")
	fmt.Println("4. 除法 (÷)")
	fmt.Println("5. 第一个复数的模长")
	fmt.Println("6. 第二个复数的模长")
	fmt.Println("7. 显示两个复数")
	fmt.Print("请输入选择 (1-7): ")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		result := c1.Add(c2)
		fmt.Printf("\n%s + %s = %s\n", c1, c2, result)

	case 2:
		result := c1.Subtract(c2)
		fmt.Printf("\n%s - %s = %s\n", c1, c2, result)

	case 3:
		result := c1.Multiply(c2)
		fmt.Printf("\n%s × %s = %s\n", c1, c2, result)

	case 4:
		if r2 == 0 && i2 == 0 {
			fmt.Println("\n错误: 除数不能为零!")
		} else {
			result := c1.Divide(c2)
			fmt.Printf("\n%s ÷ %s = %s\n", c1, c2, result)
		}

	case 5:
		mag := c1.Magnitude()
		fmt.Printf("\n|%s| = %.2f\n", c1, mag)

	case 6:
		mag := c2.Magnitude()
		fmt.Printf("\n|%s| = %.2f\n", c2, mag)

	case 7:
		fmt.Printf("\n第一个复数: %s\n", c1)
		fmt.Printf("第二个复数: %s\n", c2)

	default:
		fmt.Println("\n错误: 无效的选择! 请输入1-7之间的数字")
	}
}
