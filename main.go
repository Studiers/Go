package main

import (
	"fmt"
	//"go_study/assertion"
	"unicode"
	"go_study/item"
	a "go_study/assertion"
	_ "go_study/shape"
)

const (
	FILE_NAME = iota
	RECEIVE
	ARTORIA
)

func main() {
	//a, b := 10, 0
	arr := []rune{'a', 'b', 'c'}
	str := string(arr)
	unicode.IsDigit('1')
	fmt.Print(str)

	// Module (mylib)
	item.LibTest()

	for i := range arr {
		fmt.Println(i)
	}

	a.CheckType(int64(1))

	a := map[int]string{1:"test"}
	for k, v := range a {
		fmt.Println(k, v)
	}

	// WOW Pointer
	value := 1234
	ptrValue := &value
	ptrPtrValue := & ptrValue

	fmt.Println(value, *ptrValue, **ptrPtrValue)

	// type <custom_name> dd
	// typedef 같은 느낌?

	//r := rect {6789}
	//fmt.Println("Rect Test", r.r)
	//var c
	//fmt.Fprintf(os.Stdout, "AAA %d %d", a, b)
}
