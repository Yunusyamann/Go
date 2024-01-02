package loops

import "fmt"

//Sayı Tahmin Workshop  IF ELSE ile

//Sayı Tahmin Etme IF ELSE ile

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin tahmin ediniz?")
	fmt.Scanln(&tahminEdilenSayi)
	fmt.Println(tahminEdilenSayi)

	if aklimdakiSayi > tahminEdilenSayi {
		fmt.Println("Artımanız gerekiyor")
	} else if tahminEdilenSayi == aklimdakiSayi {
		fmt.Println("Bravo Bildiniz")
	} else if tahminEdilenSayi > aklimdakiSayi {
		fmt.Println("Hatalı")
	} else {
		fmt.Println("Tekrarr")
	}

}
