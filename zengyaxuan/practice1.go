package main

import (
	"errors"
	"fmt"
	"math"
)

type Complex struct {
	Real float64
	Imag float64
}

// 构造结构体
func NewComplex(real, imag float64) *Complex {
	return &Complex{Real: real, Imag: imag}
}

// 加法
func (c *Complex) Add(other *Complex) *Complex {
	return &Complex{
		c.Real + other.Real,
		c.Imag + other.Imag,
	}
}

// 减法
func (c *Complex) Sub(other *Complex) *Complex {
	return &Complex{
		c.Real - other.Real,
		c.Imag - other.Imag,
	}
}

// 乘法
func (c *Complex) Mul(other *Complex) *Complex {
	return &Complex{
		c.Real*other.Real - c.Imag*other.Imag,
		c.Real*other.Imag + c.Imag*other.Real,
	}
}

// 除法
func (c *Complex) Div(other *Complex) (*Complex, error) {
	divisor := other.Real*other.Real + other.Imag*other.Imag
	if divisor == 0 {
		return nil, errors.New("division by zero")
	}
	return &Complex{
		Real: (c.Real*other.Real + c.Imag*other.Imag) / divisor,
		Imag: (c.Imag*other.Real - c.Real*other.Imag) / divisor,
	}, nil
}

// 模长
func (c *Complex) Mod() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imag*c.Imag)
}

// 字符串化
func (c *Complex) String() string {
	if c.Imag > 0 {
		return fmt.Sprintf("%.2f+%.2fi", c.Real, c.Imag)
	} else if c.Imag < 0 {
		return fmt.Sprintf("%.2f%.2fi", c.Real, c.Imag)
	} else {
		return fmt.Sprintf("%.2f", c.Real)
	}
}

// 接收复数
func receiveComplex(prompt string) *Complex {
	for {
		fmt.Println(prompt + " (格式: a+bi或a-bi ):")
		var line string
		if _, err := fmt.Scanln(&line); err == nil {
			var real, imag float64
			n, err := fmt.Sscanf(line, "%f+%fi", &real, &imag)
			if err == nil && n == 2 {
				return NewComplex(real, imag)
			}
			n, err = fmt.Sscanf(line, "%f-%fi", &real, &imag)
			if err == nil && n == 2 {
				return NewComplex(real, -imag)
			}
		}
	}

}

func main() {
	c1 := receiveComplex("输入第一个复数")
	c2 := receiveComplex("输入第二个复数")

	for {
		fmt.Printf("\nc1 = %s ,  c2 = %s\n", c1, c2)
		fmt.Print("1加 2减 3乘 4除 5模长 0退出\n选择: ")
		var option int
		if _, err := fmt.Scan(&option); err != nil {
			continue
		}
		//每个选项对应的操作
		switch option {
		case 1:
			fmt.Println("c1 + c2 =", c1.Add(c2))
		case 2:
			fmt.Println("c1 - c2 =", c1.Sub(c2))
		case 3:
			fmt.Println("c1 * c2 =", c1.Mul(c2))
		case 4:
			if q, err := c1.Div(c2); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("c1 / c2 =", q)
			}
		case 5:
			fmt.Printf("|c1| = %.2f , |c2| = %.2f\n", c1.Mod(), c2.Mod())
		case 0:
			fmt.Println("bye")
			return
		default:
			fmt.Println("该选项不存在")
		}
	}
}
