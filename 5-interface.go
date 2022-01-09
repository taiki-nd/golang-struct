package main

import "fmt"

type Human interface {
	say() string
}

type Person struct {
	Name string
}

func (p *Person) say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.say() == "Mr.Mike" {
		fmt.Print("run\n")
	} else {
		fmt.Println("get out\n")　
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	//mike.say()
	DriveCar(mike)
	DriveCar(x)
}
