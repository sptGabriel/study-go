package structs

import "fmt"

type Customer struct {
	Name  string
	Email string
	CPF   int
}

func (c *Customer) Show() {
	fmt.Printf("Customer %s:\n Email: %s. Cpf: %d\n", c.Name, c.Email, c.CPF)
}
