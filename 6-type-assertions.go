package main

import "fmt"

//func do(i int) 通常は型を指定する必要がある
//
//}

func do(i interface{}) { //これでさまざまな型が使える
	/*
		ii := i.(int) //ここでtype assertion
		ii *= 2
		fmt.Println(ii)
	*/

	ss := i.(string)
	fmt.Println(ss + "!")
}

func main() {
	/*
		//do(10)
		var i interface{} = 10
		do(i) //do(10)をわかりやすく書くとこうなる
	*/

	do("Mike")
}
