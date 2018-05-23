package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{10.0,20.0}
	circle := Circle{10.2}
	fmt.Println(getArea(rect))
	fmt.Println(getArea(circle))

}
type Shape interface{
	area() float64
}
type Rectangle struct{
	width float64
	height float64
}
type Circle struct{
	radious float64
}
func (r Rectangle) area() float64{
	return r.width * r.height
}
func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radious, 2)
}
func getArea(shape Shape) float64{
	return shape.area()
}