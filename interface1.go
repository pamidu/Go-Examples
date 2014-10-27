package main
import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type square struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (s square) area()float64{
	return s.width*s.height
}
func (c circle) area()float64{
	return math.Pi*c.radius*c.radius
}

func (s square) perim()float64 {
	return (s.width+s.height)*2
	
}
func (c circle) perim()float64 {
	return math.Pi*c.radius*2
	
}
func mesuer(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {

	s:=square{width:3,height:4}
	c:=circle{radius:5}

	mesuer(s)
	mesuer(c)
}