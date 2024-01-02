package interfaces

import (
	"fmt"
	"math"
)

//Dikdortgen - Cember alan hesaplama

type shape interface { //shape-sekil
	area() float64
}

type rectyangle struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (r rectyangle) area() float64 {
	return r.height * r.widht
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şeklin Alanı :", s.area())
}

func Demo1() {
	r := rectyangle{widht: 10, height: 6}
	school(r)

	c := circle{radius: 10}
	school(c)
}
