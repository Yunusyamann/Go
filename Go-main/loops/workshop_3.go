package loops

import "fmt"

//1.Kullanıcıdan bir sayı girmesi algoritması.
//23:Asaldir

func Demo7() {

	sayi := 0
	fmt.Println("Bir Sayı Giriniz..")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}
	if asalMi == true {
		fmt.Println("Asaldır.")
	} else {
		fmt.Println("Asal degildir.")
	}

}
