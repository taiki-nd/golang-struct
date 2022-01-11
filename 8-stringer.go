package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) String() string { //stringのメソッド実装すれば出力結果が変わる
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
