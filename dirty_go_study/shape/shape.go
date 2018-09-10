package shape

import "fmt"

type shape interface {
	area() float64
}

func Describe(s shape) {
	fmt.Println(s.area())
}
