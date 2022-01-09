package main

import "fmt"

/*
type Vertex struct {
	X, Y int
}

func area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(area(v))//vとarea関数に結びつきがない
}
*/

//メソッド（.で呼び出す）
/*
type Vertex struct {
	X, Y int
}

func (v Vertex) area() int { //値レシーバー
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.area())
}
*/

type Vertex struct {
	X, Y int
}

func (v Vertex) area() int {
	return v.X * v.Y
}

func (v *Vertex) scale(i int) { //ポイントレシーバー
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.area())
	v.scale(10)
	fmt.Println(v.area())
}
