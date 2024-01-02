package functions

import "fmt"


func Topla(sayi1 int, sayi2 int) {
	var toplam = sayi1 + sayi2
	fmt.Println("Sonuç :", toplam)
}

func Toplam(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	return toplam
}

func SelamVer(kullaniciAdi string) {
	fmt.Println("Hello ", kullaniciAdi)
}
