package main

import "fmt"

type Vertex struct {
	x, y int
}

/*
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
*/
type Vertex3D struct {
	Vertex //x, y intがtype Vertexによって定義されている
	z      int
}

func (v Vertex3D) area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	/*
		v := Vertex{3, 4}
		v.scale(10)
		fmt.Println(v.area())

		vv := New(3, 4)
		vv.scale(10)
		fmt.Println(vv.area())
	*/

	vvv := Vertex3D{Vertex{3, 4}, 5}
	vvv.scale3D(10)
	fmt.Println(vvv.area3D())

	vvvv := New3D(3, 4, 5)
	vvvv.scale3D(10)
	fmt.Println(vvvv.area3D())
}
