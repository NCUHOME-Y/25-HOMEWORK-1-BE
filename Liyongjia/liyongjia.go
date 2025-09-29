package main

import (
	"fmt"
	"math"
)

// 复数结构体
type Complex struct {
	real float64
	imag float64
}

// https://blog.csdn.net/weixin_42004975/article/details/111589053 公式链接

// 加法
func (c Complex) Add(other Complex) Complex {
	return Complex{
		real: c.real + other.real,
		imag: c.imag + other.imag,
	}
}

// 减法
func (c Complex) Sub(other Complex) Complex {
	return Complex{
		real: c.real - other.real,
		imag: c.imag - other.imag,
	}
}

// 乘法
func (c Complex) Mul(other Complex) Complex {
	return Complex{
		real: c.real*other.real - c.imag*other.imag,
		imag: c.real*other.imag + c.imag*other.real,
	}
}

// 除法
var ok = 0

func (c Complex) Div(other Complex) Complex {

	judge := other.real*other.real + other.imag*other.imag
	if judge == 0 {
		fmt.Println("除数不能为0")
		ok = 1
	}
	return Complex{
		real: (c.real*other.real + c.imag*other.imag) / (other.real*other.real + other.imag*other.imag),
		imag: (c.imag*other.real - c.real*other.imag) / (other.real*other.real + other.imag*other.imag),
	}
}

// 模长
func (c Complex) Len() float64 {
	return math.Sqrt(c.real*c.real + c.imag*c.imag) //开平方根
}

// 转化为 字符串
func (c Complex) ToString() string {
	if c.imag > 0 {
		return fmt.Sprintf("%.2f + %.2fi", c.real, c.imag)
	} else if c.imag == 0 {
		return fmt.Sprintf("%.2f", c.real)
	} else if c.imag < 0 {
		return fmt.Sprintf("%.2f - %.2fi", c.real, -c.imag)
	}
	return "wrong"
}
func main() {
	// // 测试
	// c1 := Complex{real: 1, imag: 2}
	// c2 := Complex{real: 3, imag: 4}

	var real1, imag1, real2, imag2 float64
	fmt.Println("请输入第一个复数的实部和虚部：")
	fmt.Scan(&real1, &imag1)
	fmt.Println("请输入第二个复数的实部和虚部：")
	fmt.Scan(&real2, &imag2)
	c1 := Complex{real: real1, imag: imag1}
	c2 := Complex{real: real2, imag: imag2}

	//用虚部的正负号来决定输出的格式  if

	c3 := c1.Add(c2) //加法
	if c3.imag >= 0 {
		fmt.Printf("这是加法 c1 + c2 = %v + %vi\n", c3.real, c3.imag)
	} else {
		fmt.Printf("这是加法 c1 + c2 = %v - %vi\n", c3.real, -c3.imag)
	}

	c4 := c1.Sub(c2) //减法
	if c4.imag >= 0 {
		fmt.Printf("这是减法 c1 - c2 = %v + %vi\n", c4.real, c4.imag)
	} else {
		fmt.Printf("这是减法 c1 - c2 = %v - %vi\n", c4.real, -c4.imag)
	}

	c5 := c1.Mul(c2) //乘法
	if c5.imag >= 0 {
		fmt.Printf("这是乘法 c1 * c2 = %v + %vi\n", c5.real, c5.imag)
	} else {
		fmt.Printf("这是乘法 c1 * c2 = %v - %vi\n", c5.real, -c5.imag)
	}

	c6 := c1.Div(c2) //除法
	if ok == 0 {
		if c6.imag >= 0 {
			fmt.Printf("这是除法 c1 / c2 = %.2f + %.2fi\n", c6.real, c6.imag)
		} else {
			fmt.Printf("这是除法 c1 / c2 = %.2f - %.2fi\n", c6.real, -c6.imag)
		}
	}
	c7 := c1.Len() //模长
	fmt.Printf("这是模长 |c1| = %.2f\n", c7)

	c8 := c1.ToString() //字符串
	fmt.Printf("这是字符串 c1 = %s\n", c8)

}
