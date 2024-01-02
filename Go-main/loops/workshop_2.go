package loops

import "fmt"

//Sayı tahmininde kaçıncı adımda doğru tahmin yapıldı Algoritması.
//1-3:Super  4-6:Guzel >6 Fena Degil

func Demo6() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir sayı tahmin tahmin ediniz?")
	fmt.Scanln(&tahminEdilenSayi) //Kullanıcıdan Bilgi alma
	tahminSayisi += 1

	for aklimdakiSayi != tahminEdilenSayi { //aklimdakiSayi tahminEdilenSayi farkli oldugu surece devam etmesi.
		if aklimdakiSayi > tahminEdilenSayi {
			fmt.Println("Daha büyük sayı giriniz..")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi += 1
		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha küçük sayı giriniz..")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi += 1
		}
	}

	basariDurumu := ""

	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Super"
	} else if tahminSayisi < 6 {
		basariDurumu = "Guzel"
	} else {
		basariDurumu = "Fena Degil "
	}
	fmt.Printf("Bravo Bildiniz.  %v tahmin  %v", tahminSayisi, basariDurumu)
}
