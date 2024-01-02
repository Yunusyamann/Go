package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi :", c.firstName, c.lastName, c.age)
}

func (c customer) update() {
	fmt.Println("Güncellendi :", c.firstName, c.lastName, c.age)
}

func Demo2() {
	c := customer{firstName: "Ozcan", lastName: "Kara", age: 23}
	c.save()
	c.update()
}
