package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sptGabriel/study-go/maps"
	"github.com/sptGabriel/study-go/math"
	"github.com/sptGabriel/study-go/pointers"
	"github.com/sptGabriel/study-go/structs"
)

func main() {
	result := math.SumX(1)
	fmt.Printf("Sum: %v \n", result)
	fmt.Printf("Sum: %v \n", result)
	fmt.Println("pointers:")
	aPointer := 10
	pointers.PointerExample(&aPointer)
	fmt.Printf("Sum: %v \n", aPointer)
	fmt.Println("pointers / struct:")
	person := pointers.Person{
		FirstName: "Biel",
		LastName:  "Costa",
	}
	person.FullName()
	fmt.Println("Structs:")
	customer := structs.Customer{
		Name:  "Gabriel",
		Email: "spt@spt.com",
		CPF:   12345678,
	}
	customer2 := structs.Customer{"xd", "xd@spt.com", 455445342}
	customer2.Show()
	customer.Show()
	fmt.Println("Convert to JSON:")
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Marshal bytes slice", customerJSON)
	fmt.Println(string(customerJSON))
	fmt.Println("Convert to Struct:")
	jsonCustomer3 := `{"Name":"Gabriel","Email":"spt@spt.com","CPF":12345678}`
	customer3 := structs.Customer{}
	json.Unmarshal([]byte(jsonCustomer3), &customer3)
	fmt.Println(customer3)
	fmt.Println("GO MAPS:")
	maps.Test()
	maps.SyncMAP()
}
