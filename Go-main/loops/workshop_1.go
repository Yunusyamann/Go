package loops

import "fmt"

//Sayı Tahmin Etme For döngüsü ile
func Demo4() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	fmt.Println("Bir sayı tahmin tahmin ediniz?")
	for i := 0; tahminEdilenSayi < aklimdakiSayi; i++ {
		fmt.Scanln(&tahminEdilenSayi)
	}
	fmt.Println(tahminEdilenSayi)

	fmt.Println("-----------------------------------")

}
