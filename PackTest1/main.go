package main

import (
	"PackTest1/computation"
	"fmt"
)

var (
	a int = 2
	b int = 3
)

func main() {
	fmt.Println("Hello")
	fmt.Println(computation.Add1(a, b))
	fmt.Println(computation.Sub1(a, b))

}
