package main

import "fmt"

type Vertex struct {
	x, y int
}

func area(v Vertex) int {
	return v.x * v.y
}

func (v Vertex) area() int {
	return v.x * v.y
}

func (v *Vertex) scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.area())

	vv := New(3, 4)
	vv.scale(10)
	fmt.Println(vv.area())
}
