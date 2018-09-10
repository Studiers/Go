package assertion

import "fmt"

func CheckType(v interface{}) {
	switch v.(type) {
	case int64:
		fmt.Println("Int64!!")
	default:
		fmt.Println("It is not Int64!!")
	}
}