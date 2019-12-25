package main

import "fmt"

type A struct {
	B
}

type B struct {

}

func (b *B)b()  {
	fmt.Println("b")
}

func main() {
	a := A{}
	a.b()
}