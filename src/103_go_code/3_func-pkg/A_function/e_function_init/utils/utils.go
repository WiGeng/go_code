package utils

import "fmt"

var Car string
var Color string

func init() {
	fmt.Println("utils中的init()~")
	Car = "Benz"
	Color = "Black"
}
