package conditionals

import "fmt"

func Demo4() {
	//3 adet int değişken tanmımlama
	//Ekrana en büyük olanı yazdırma.

	var sayi1, sayi2, sayi3 int = 10, 5, 18
	var enBuyuk int = sayi1
	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	} else if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}
	fmt.Printf("En büyük sayı : %v", enBuyuk)
}
