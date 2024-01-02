package conditionals

import "fmt"

func Demo3() {
	//3 adet int değişken tanmımlama
	//Ekrana en büyük olanı yazdırma.

	var sayi1 int = 620
	var sayi2 int = 200
	var sayi3 int = 5098

	if sayi1 > sayi2 {
		fmt.Printf("En büyük Sayı :sayi1 ")
	} else if sayi1 > sayi3 {
		fmt.Printf("En büyük Sayı :sayi1 ")
	} else if sayi2 > sayi1 {
		fmt.Printf("En büyük Sayı :sayi2")
	} else if sayi2 > sayi3 {
		fmt.Printf("En büyük Sayı :sayi2")
	} else if sayi3 > sayi1 {
		fmt.Printf("En büyük Sayı :sayi3")
	} else if sayi3 > sayi2 {
		fmt.Printf("En büyük Sayı :sayi3")
	} else {
		fmt.Printf("Tekrar Deneyiniz ...")
	}
}
