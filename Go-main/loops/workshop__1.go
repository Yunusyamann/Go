package loops

import "fmt"

//Sayı Tahmin Etme If Else -For  döngüsü ile
func Demo5() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin tahmin ediniz?")
	fmt.Scanln(&tahminEdilenSayi) //Kullanıcıdan Bilgi alma

	for aklimdakiSayi != tahminEdilenSayi { //aklimdakiSayi tahminEdilenSayi farkli oldugu surece devam etmesi.
		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Daha büyük sayı giriniz..")
			fmt.Scanln(&tahminEdilenSayi)
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük sayı giriniz..")
			fmt.Scanln(&tahminEdilenSayi)
		}
	}
	fmt.Println("Bravo Bildiniz")
}
