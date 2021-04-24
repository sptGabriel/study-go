package pointers

import "fmt"

func PointerExample(value *int) {
	fmt.Println("this function go set value of a pointer")
	*value = 20
}
