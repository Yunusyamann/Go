package error_handling

import (
	"errors"
	"fmt"
)

//Errors_Librariy

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 55

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasinda bir sayi giriniz.")
	}

	if tahmin > aklimdakiSayi {
		return "Daha kucuk sayi giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha buyuk sayi giriniz", nil
	}
	return "Bildiniz", nil
}
func Demo3() {
	mesaj, _ := TahminEt(51)
	fmt.Println(mesaj)

	mesaj2, hata := TahminEt(49)
	fmt.Println(mesaj2)
	fmt.Println(hata)

	//fmt.Println(TahminEt(101))
	//fmt.Println(TahminEt(-100))
}
