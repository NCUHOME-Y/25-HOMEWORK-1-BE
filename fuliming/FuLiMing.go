package main

import (
	"fmt"
	"math"
	"errors"
)

type Complex struct {
	real float64
	fade float64
}

var a, b, c, d float64

func newComplex(Real, Fade float64) *Complex {
	return &Complex{real: Real, fade: Fade}
}

func (one *Complex) Add(another *Complex) *Complex {
	return &Complex{
		real: one.real + another.real,
		fade: one.fade + another.fade,
	}
}

func (one *Complex) Sub(another *Complex) *Complex {
	return &Complex{
		real: one.real - another.real,
		fade: one.fade - another.fade,
	}
}

func (one *Complex) Mul(another *Complex) *Complex {
	return &Complex{
		real: one.real * another.real - one.fade * another.fade,
		fade: one.fade * another.real + one.real * another.fade,
}
}

func (one *Complex) Div(another *Complex) (*Complex, error) {
    denominator := another.real*another.real + another.fade*another.fade
	    if denominator == 0 {
		    return nil, errors.New("division by zero")
	    }
	    return &Complex{
		    real: (one.real*another.real + one.fade*another.fade) / denominator,
		    fade: (one.fade*another.real - one.real*another.fade) / denominator,
	}, nil
}

func (one *Complex) Abs() float64 {
	return math.Sqrt(one.real * one.real + one.fade * one.fade)
}
	



func (one *Complex) ToString() string {
	if one.fade == 0 {
		return fmt.Sprintf("%g", one.real)
	}

	if one.real == 0 {
		return fmt.Sprintf("%gi", one.fade)
	}
	
	if one.fade > 0 {
		return fmt.Sprintf("%g+%gi", one.real, one.fade)
	}
	return fmt.Sprintf("%g%gi", one.real, one.fade) 
}

func main() {
	fmt.Printf("请输入第一个复数的实部与虚部（用空格隔开）：")
	fmt.Scanln(&a, &b)
	fmt.Printf("请输入另一个复数的实部与虚部（用空格隔开）：")
	fmt.Scanln(&c, &d)

	comPlex1 := newComplex(a, b)
	comPlex2 := newComplex(c, d)

	fmt.Printf("comPlex1 = %s\n", comPlex1.ToString())
	fmt.Printf("comPlex2 = %s\n", comPlex2.ToString())
	fmt.Printf("模长 |comPlex1| = %.2f\n", comPlex1.Abs())

        sum := comPlex1.Add(comPlex2)
	    fmt.Printf("加法: %s + %s = %s\n", comPlex1.ToString(), comPlex2.ToString(), sum.ToString())
	
	
	    diff := comPlex1.Sub(comPlex2)
	    fmt.Printf("减法: %s - %s = %s\n", comPlex1.ToString(), comPlex2.ToString(), diff.ToString())
	
	
	    product := comPlex1.Mul(comPlex2)
	    fmt.Printf("乘法: %s × %s = %s\n", comPlex1.ToString(), comPlex2.ToString(), product.ToString())
	

	    quotient, err := comPlex1.Div(comPlex2)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法: %s / %s = %s\n", comPlex1.ToString(), comPlex2.ToString(), quotient.ToString())
	}
}