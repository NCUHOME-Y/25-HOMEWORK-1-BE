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

// 构造结构体的函数
func NewComplex(real, imag float64) Complex {
	return Complex{
		Real: real,
		Imag: imag,
	}
}

// 加法函数
func (count1 Complex) Add(count2 Complex) Complex {
	return Complex{
		Real: count1.Real + count2.Real,
		Imag: count1.Imag + count2.Imag,
	}
}

// 减法函数
func (count1 Complex) Subtract(count2 Complex) Complex {
	return Complex{
		Real: count1.Real - count2.Real,
		Imag: count1.Imag - count2.Imag,
	}
}

// 乘法函数
func (count1 Complex) Multiply(count2 Complex) Complex {
	realPart := count1.Real*count2.Real - count1.Imag*count2.Imag
	imagPart := count1.Real*count2.Imag + count1.Imag*count2.Real
	return Complex{
		Real: realPart,
		Imag: imagPart,
	}
}

// 除法函数
func (count1 Complex) Divide(count2 Complex) (Complex, error) {
	mother := count2.Real*count2.Real + count2.Imag*count2.Imag
	if mother == 0 {
		return Complex{}, errors.New("错误：分母为零")
	} else {
		realSon := count1.Real*count2.Real + count1.Imag*count2.Imag
		imagSon := count1.Imag*count2.Real - count1.Real*count2.Imag
		return Complex{
			Real: realSon / mother,
			Imag: imagSon / mother,
		}, nil
	}
}

// Magnitude 实现求复数的模长 (绝对值)：|a+bi| = (a^2 + b^2)开根
func (count3 Complex) Magnitude() float64 {
	return math.Hypot(count3.Real, count3.Imag)
}

// 打印运算后结果的函数（寻求AI帮助）
func (count3 Complex) ToString() string {
	if count3.Imag == 0 {
		return fmt.Sprintf("%.2f", count3.Real)
	}
	if count3.Real == 0 {
		// 纯虚数
		return fmt.Sprintf("%.2fi", count3.Imag)
	}
	if count3.Imag < 0 {
		// 负虚部：Real - |Imag|i
		return fmt.Sprintf("%.2f - %.2fi", count3.Real, math.Abs(count3.Imag))
	}
	return fmt.Sprintf("%.2f + %.2fi", count3.Real, count3.Imag)
}

// 获取用户输入的复数
func getUserComplexInput(name string) Complex {
	var realPart, imagPart float64

	fmt.Printf("请输入复数 %s 的实部 (Real): ", name)
	_, err := fmt.Scanf("%f\n", &realPart)
	if err != nil {
		fmt.Println("输入错误，使用默认值 0.0")
		realPart = 0.0
	}

	fmt.Printf("请输入复数 %s 的虚部 (Imag): ", name)
	_, err = fmt.Scanf("%f\n", &imagPart)
	if err != nil {
		fmt.Println("输入错误，使用默认值 0.0")
		imagPart = 0.0
	}

	return NewComplex(realPart, imagPart)
}

func main() {
	// 1. 获取用户输入：复数 Z1
	z1 := getUserComplexInput("Z1")
	fmt.Printf("Z1 : %s\n", z1.ToString())

	// 2. 获取用户输入：复数 Z2
	z2 := getUserComplexInput("Z2")
	fmt.Printf("Z2 : %s\n", z2.ToString())

	fmt.Println("请选择您想进行的操作:")
	fmt.Println(" 1: Z1 + Z2 (加法)")
	fmt.Println(" 2: Z1 - Z2 (减法)")
	fmt.Println(" 3: Z1 * Z2 (乘法)")
	fmt.Println(" 4: Z1 / Z2 (除法)")
	fmt.Println(" 5: |Z1| (Z1 的模长)")
	fmt.Println(" Q: 退出程序")

	var choice string

	// 使用 for 循环创建一个持续运行的菜单，直到用户选择退出
	for {
		fmt.Print("请输入您的选择 (1-5 或 Q): ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("输入读取失败，请重试。", err)
			continue
		}

		// 使用 switch 结构根据用户输入调用对应的方法
		switch choice {
		case "1":
			result := z1.Add(z2)
			fmt.Printf("结果 Z1 + Z2 = %s\n", result.ToString())

		case "2":
			result := z1.Subtract(z2)
			fmt.Printf("结果 Z1 - Z2 = %s\n", result.ToString())

		case "3":
			result := z1.Multiply(z2)
			fmt.Printf("结果 Z1 * Z2 = %s\n", result.ToString())

		case "4":
			result, err := z1.Divide(z2)
			if err != nil {
				fmt.Printf("除法错误: %v\n", err)
			} else {
				fmt.Printf("结果 Z1 / Z2 = %s\n", result.ToString())
			}

		case "5":
			// 模长（问了AI，不清楚模长怎么求，嘿嘿）
			result := z1.Magnitude()
			fmt.Printf("结果 |Z1| = %.2f\n", result)

		case "Q":
			// 退出程序
			fmt.Println("白白咯！")
			return
		default:
			fmt.Println("无效的选项，请重新输入。")
		}
	}
}
