package pointers

import "fmt"

//Pointer :Bellekteki degişkenlerin tutuldugu yer.

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı :", *sayi)
}
