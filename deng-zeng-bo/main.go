package main

import (
	"fmt"
	"math"
)

// Complex 自定义复数类型
type Complex struct {
	Real float64 // 实部
	Imag float64 // 虚部
}

// Add 加法：(a+bi)+(c+di) = (a+c)+(b+d)i
func (c Complex) Add(other Complex) Complex {
	return Complex{
		Real: c.Real + other.Real,
		Imag: c.Imag + other.Imag,
	}
}

// Sub 减法：(a+bi)-(c+di) = (a-c)+(b-d)i
func (c Complex) Sub(other Complex) Complex {
	return Complex{
		Real: c.Real - other.Real,
		Imag: c.Imag - other.Imag,
	}
}

// Mul 乘法：(a+bi)*(c+di) = (ac - bd)+(ad + bc)i
func (c Complex) Mul(other Complex) Complex {
	return Complex{
		Real: c.Real*other.Real - c.Imag*other.Imag,
		Imag: c.Real*other.Imag + c.Imag*other.Real,
	}
}

// Div 除法：(a+bi)/(c+di) = [(ac + bd)+(bc - ad)i] / (c² + d²)
func (c Complex) Div(other Complex) Complex {
	denominator := other.Real*other.Real + other.Imag*other.Imag // 定义分母,简化了代码
	if denominator == 0 {
		panic("除数不能为零")
	}
	return Complex{
		Real: (c.Real*other.Real + c.Imag*other.Imag) / denominator,
		Imag: (c.Imag*other.Real - c.Real*other.Imag) / denominator,
	}
}

// Modulus 求模长：Sprt(a² + b²)
func (c Complex) Modulus() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imag*c.Imag)
}

// String 实现 fmt包中Stringer 接口
// 当使用 fmt 包的打印函数（如 fmt.Println、fmt.Printf 等）输出 Complex 类型的变量时，会自动调用这个方法来获取字符串表示形式
// 统一格式化输出，简化了代码，(本段学习了ai的处理办法)
func (c Complex) String() string {
	if c.Imag == 0 {
		return fmt.Sprintf("%.2f",c.Real)
	}
	if c.Imag > 0 {
		return fmt.Sprintf("%.2f+%.2fi", c.Real, c.Imag)
	}
	return fmt.Sprintf("%.2f%.2fi", c.Real, c.Imag)
}

func main() {
	var c1, c2 Complex
	fmt.Println("请输入两个数的实部和虚部")
	fmt.Scan(&c1.Real, &c1.Imag, &c2.Real, &c2.Imag)


	// 和
	sum := c1.Add(c2)
	fmt.Println("c1 + c2 =", sum)
	// 差
	diff := c1.Sub(c2)
	fmt.Println("c1 - c2 =", diff)
	// 积
	product := c1.Mul(c2)
	fmt.Println("c1 * c2 =", product)
	// 商
	quotient := c1.Div(c2)
	fmt.Println("c1 / c2 =", quotient)
	// 模
	modulus := c1.Modulus()
	fmt.Printf("|c1| = %.2f\n", modulus)
}