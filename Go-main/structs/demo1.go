package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Laptop", 23000, "MacBook"})
	fmt.Println(product{"Laptop", 23000, ""})
	fmt.Println(product{name: "Laptop", unitPrice: 23000})
	fmt.Println(product{name: "Laptop"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
	//discountRate int
}