package main;
import(
	"fmt"
	"math"
)
type Complex struct{
	Real float64
	Imaginary float64

}
func(c Complex) Add(d Complex)Complex{
	return Complex{
		Real:c.Real+d.Real,
		Imaginary:c.Imaginary+d.Imaginary,
	}
}
func(c Complex) Subtract(d Complex)Complex{
	return Complex{
		Real:c.Real-d.Real,
		Imaginary:c.Imaginary-d.Imaginary,
	}
}
func(c Complex) Multiple(d Complex)Complex{
	return Complex{
		Real:c.Real*d.Real-c.Imaginary*d.Imaginary,
		Imaginary:c.Real*d.Imaginary + c.Imaginary*d.Real,
	}
}
func(c Complex) Divide(d Complex)Complex{
	denominator :=d.Real*d.Real + d.Imaginary*d.Imaginary
	if denominator == 0 {
		panic("除数不可以为零！笨蛋！")
	}
    return Complex{
		Real: (c.Real*d.Real + c.Imaginary*d.Imaginary) / denominator,
		Imaginary: (c.Imaginary*d.Real - c.Real*d.Imaginary) / denominator,
	}
}
func(c Complex)Magnitude()float64{
	return math.Sqrt(c.Real*c.Real + c.Imaginary*c.Imaginary)
}
func (c Complex) String() string {
	if c.Imaginary < 0 {
		return fmt.Sprintf("%.2f - %.2fi", c.Real, math.Abs(c.Imaginary))//取绝对值，否则他会输出一个+-出来吓晕了
	}
	return fmt.Sprintf("%.2f + %.2fi", c.Real, c.Imaginary)
}
func main(){
	c1 := Complex{Real: 3, Imaginary: 4}
	c2 := Complex{Real: 5, Imaginary: -12}
	fmt.Printf("c1 = %v\n", c1)
	fmt.Printf("c2 = %v\n",c2)
	//加法
	sum := c1.Add(c2)
	fmt.Printf("c1 + c2 = %v\n", sum)
	//减法
	diff:=c1.Subtract(c2)
	fmt.Printf("c1-c2=%v\n",diff)
	//乘法
	prod:=c1.Multiple(c2)
	fmt.Printf("c1*c2=%v\n",prod)
	//除法
	div:=c1.Divide(c2)
	fmt.Printf("c1/c2=%v\n",div)
	//取模
	module1:=c1.Magnitude()
	module2:=c2.Magnitude()
	fmt.Printf("|c1|=%2f\n",module1)
	fmt.Printf("|c2|=%2f\n",module2)


}
