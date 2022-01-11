/*
Q2 X is 3! Y is 4! と表示されるStringerを作成してください。

package main

import (
    "fmt"
)

type Vertex struct{
    X, Y int
}

func main(){
    v := Vertex{3, 4}
    fmt.Println(v)
}
*/

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v, Y is %v", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
