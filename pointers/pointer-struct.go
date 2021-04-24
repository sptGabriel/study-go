package pointers

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) FullName() {
	p.FirstName = "Gabriel"
	fmt.Println("Full Name:", p.FirstName, p.LastName)
}
