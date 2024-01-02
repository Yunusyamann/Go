package slices

import "fmt"

// Slice Gelismis Dizi
func Demo1() {
	isimler := make([]string, 3) //make ile olusturma

	fmt.Println(isimler)

	isimler[0] = "Yunus"
	isimler[1] = "Derin"
	isimler[2] = "Salih"
	isimler = append(isimler, "Berfin") //append ile Yeni Eleman ekleme

	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
