package interfaces

import "fmt"

//type assertion mantigi
func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "Ozcan"
	dogrula(sayi1)

	var sayi2 interface{} = 100
	dogrula(sayi2)
}