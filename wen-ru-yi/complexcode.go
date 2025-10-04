package main

import (
	"fmt"
	"math"
)

// 加法
func addComplex(r1, i1, r2, i2 float64) (float64, float64) {
	return r1 + r2, i1 + i2
}

// 减法
func subtractComplex(r1, i1, r2, i2 float64) (float64, float64) {
	return r1 - r2, i1 - i2
}

// 乘法
func multiplyComplex(r1, i1, r2, i2 float64) (float64, float64) {
	realPart := r1*r2 - i1*i2
	imagPart := r1*i2 + i1*r2
	return realPart, imagPart
}

// 除法
func divideComplex(r1, i1, r2, i2 float64) (float64, float64) {
	denominator := r2*r2 + i2*i2
	realPart := (r1*r2 + i1*i2) / denominator
	imagPart := (i1*r2 - r1*i2) / denominator
	return realPart, imagPart
}

// 模长
func magnitudeComplex(r, i float64) float64 {
	return math.Sqrt(r*r + i*i)
}

// 显示复数
func showComplex(r, i float64) string {
	return fmt.Sprintf("(%g + %gi)", r, i)
}

func main() {
	var choice int
	var r1, i1, r2, i2 float64

	fmt.Println("=== 复数计算器 ===")

	// 输入第一个复数
	fmt.Print("请输入第一个复数的实部: ")
	fmt.Scan(&r1)
	fmt.Print("请输入第一个复数的虚部: ")
	fmt.Scan(&i1)

	// 输入第二个复数
	fmt.Print("请输入第二个复数的实部: ")
	fmt.Scan(&r2)
	fmt.Print("请输入第二个复数的虚部: ")
	fmt.Scan(&i2)

	fmt.Printf("\n第一个复数: %s\n", showComplex(r1, i1))
	fmt.Printf("第二个复数: %s\n", showComplex(r2, i2))

	// 显示运算菜单
	fmt.Println("\n请选择运算:")
	fmt.Println("1. 加法 (+)")
	fmt.Println("2. 减法 (-)")
	fmt.Println("3. 乘法 (×)")
	fmt.Println("4. 除法 (÷)")
	fmt.Println("5. 第一个复数的模长")
	fmt.Println("6. 第二个复数的模长")
	fmt.Print("请输入选择 (1-6): ")

	fmt.Scan(&choice)

	// 使用switch进行运算选择
	switch choice {
	case 1:
		rResult, iResult := addComplex(r1, i1, r2, i2)
		fmt.Printf("\n结果: %s + %s = %s\n",
			showComplex(r1, i1), showComplex(r2, i2), showComplex(rResult, iResult))

	case 2:
		rResult, iResult := subtractComplex(r1, i1, r2, i2)
		fmt.Printf("\n结果: %s - %s = %s\n",
			showComplex(r1, i1), showComplex(r2, i2), showComplex(rResult, iResult))

	case 3:
		rResult, iResult := multiplyComplex(r1, i1, r2, i2)
		fmt.Printf("\n结果: %s × %s = %s\n",
			showComplex(r1, i1), showComplex(r2, i2), showComplex(rResult, iResult))

	case 4:
		if r2 == 0 && i2 == 0 {
			fmt.Println("\n错误: 除数不能为零!")
		} else {
			rResult, iResult := divideComplex(r1, i1, r2, i2)
			fmt.Printf("\n结果: %s ÷ %s = %s\n",
				showComplex(r1, i1), showComplex(r2, i2), showComplex(rResult, iResult))
		}

	case 5:
		mag := magnitudeComplex(r1, i1)
		fmt.Printf("\n结果: |%s| = %g\n", showComplex(r1, i1), mag)

	case 6:
		mag := magnitudeComplex(r2, i2)
		fmt.Printf("\n结果: |%s| = %g\n", showComplex(r2, i2), mag)

	default:
		fmt.Println("\n错误: 无效的选择! 请输入1-6之间的数字")
	}
}
