package main

import "fmt"

type rect struct{
	width , height int
}

func (r *rect) area() int {
	return r.width*r.height
}

func (r rect) perim() int{
	return 2*(r.width+r.height)
}



func main(){
	r:= rect{width: 20,height:30}
	fmt.Println("Area of this square is :",r.area())
	fmt.Println("Perimeter of this Square is :",r.perim())

	rp:=&r
	fmt.Println("Area of this square is :",rp.area())
	fmt.Println("Perimeter of this Square is :",rp.perim())
	
}