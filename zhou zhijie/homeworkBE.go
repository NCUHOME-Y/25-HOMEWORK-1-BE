package main

import (
	"fmt"
	"math"
)

type Complex struct{
	Real float64
	Image float64
}

//加法
func (a *Complex) Add(b *Complex) *Complex {
	return &Complex{a.Real + b.Real, a.Image + b.Image}
}

//减法
func (a *Complex) Sub(b *Complex) *Complex {
	return &Complex{a.Real - b.Real, a.Image - b.Image}
}

//乘法
func (a *Complex) Mul(b *Complex) *Complex {
	real := a.Real*b.Real - a.Image*b.Image
	image := a.Real*b.Image + a.Image*b.Real
	return &Complex{real, image}
}

//除法
func (a *Complex) Div(b *Complex) (*Complex,error) {
	denominator := b.Real*b.Real + b.Image*b.Image
	if denominator == 0{
		return nil, fmt.Errorf("分母不应为0!!!!")
	}
	real := (a.Real*b.Real + a.Image*b.Image) / denominator
	image := (a.Image*b.Real - a.Real*b.Image) / denominator
	return &Complex{real,image} , nil
}

//模长
func (a *Complex) Mod() float64 {
	return math.Sqrt(a.Real*a.Real + a.Image*a.Image)
}

//Tostring
func (a *Complex) ToString() string{
	switch{
	case a.Image == 0:
		return fmt.Sprintf("%.2f", a.Real)
	case a.Image > 0:
		return fmt.Sprintf("%.2f+%.2fi",a.Real, a.Image)
	default:
		return fmt.Sprintf("%.2f+%.2fi",a.Real, -a.Image)
	}
}

//实验
func main() {
	c1 := &Complex{2, 3} // 复数 2+3i
	c2 := &Complex{1, 2} // 复数 1+2i

	fmt.Printf("c1 = %s, c2 = %s\n", c1.ToString(), c2.ToString())
	fmt.Printf("c1 + c2 = %s\n", c1.Add(c2).ToString())  // 3.00+5.00i
	fmt.Printf("c1 - c2 = %s\n", c1.Sub(c2).ToString())  // 1.00+1.00i
	fmt.Printf("c1 × c2 = %s\n", c1.Mul(c2).ToString())  // -4.00+7.00i
	if divRes, err := c1.Div(c2); err == nil {
		fmt.Printf("c1 ÷ c2 = %s\n", divRes.ToString()) // 1.60-0.20i
	}
	fmt.Printf("c1 模长 = %.4f\n", c1.Mod())         // 3.6056
}