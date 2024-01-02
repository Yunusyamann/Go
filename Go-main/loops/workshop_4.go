package loops

import "fmt"

//220 ve 284 arkadas sayilardir.
//10 ve 65 arkadas sayi degildir.
func Demo8() {
	sayi1 := 220
	sayi2 := 284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 += i
		}
	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadas sayılardır.", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadas sayılar değildir.", sayi1, sayi2)
	}

}
