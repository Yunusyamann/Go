package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{20, 30, 40, 50, 60}
	fmt.Println(sayilar)

	fmt.Println()

	//Max sayi bulma
	sayilar2 := [5]int{20, 30, 40, 50, 60}
	fmt.Println(sayilar2)

	enBuyuk := sayilar2[0]

	//lenght = uzunluk
	for i := 0; i < len(sayilar2); i++ {
		if sayilar2[i] > enBuyuk {
			enBuyuk = sayilar2[i]
		}
	}
	fmt.Println("En Büyük Sayı : ", enBuyuk)
}
