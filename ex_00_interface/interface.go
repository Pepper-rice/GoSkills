package main

import (
	"fmt"
	"math"
)

// 接口，可以理解为方法的集合
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 在 GO 中要实现一个接口，只要实现接口中的所有方法
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// 如果一个变量实现了某一个接口，我们就可以调用指定接口的方法
func area(g geometry) float64 {
	return g.area()
}

func perim(g geometry) float64 {
	return g.perim()
}

func main() {
	r := rect{width: 10, height: 20}
	c := circle{radius: 5}

	fmt.Println(area(r))
	fmt.Println(area(c))
	fmt.Println(perim(r))
	fmt.Println(perim(c))

}
