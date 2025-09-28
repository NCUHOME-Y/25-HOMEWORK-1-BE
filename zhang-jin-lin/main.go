package main

type Complex struct {
	real      float64
	imaginary float64
}

func (first Complex) Add(second Complex) Complex {
	return Complex{
		real:      first.real + second.real,
		imaginary: first.imaginary + second.imaginary,
	}
}
func (first Complex) Sub(second Complex) Complex {
	return Complex{
		real:      first.real - second.real,
		imaginary: first.imaginary - second.imaginary,
	}
}
func (first Complex) Mul(second Complex) Complex {
	return Complex{
		real:      first.real * second.real,
		imaginary: first.imaginary * second.imaginary,
	}
}
func (first Complex) Div(second Complex) Complex {
	return Complex{
		real:      first.real / second.real,
		imaginary: first.imaginary / second.imaginary,
	}
}

// 公式参考牛顿迭代求平方根//
func ModuleLen(s Complex) float64 {
	a := s.real * s.real
	b := s.imaginary * s.imaginary
	n := a*a + b*b
	if n == 0 {
		return 0
	}
	k := n
	var x float64
	for !(n-1e-10 <= x*x && x*x <= n+1e-10) {
		x = (k + n/k) / 2
		k = x
	}
	return x
}
