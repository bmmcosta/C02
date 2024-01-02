package main

import (
	"fmt"
)

var Name string
var Address string

func main() {
	Name = "Jane Doe"
	Address = "Av. Foo 123"
	fmt.Printf("%s lives at %s. \n", Name, Address)

}
