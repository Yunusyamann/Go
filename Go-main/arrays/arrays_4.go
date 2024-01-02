package arrays

import "fmt"

//Çok Boyutlu Diziler

func Demo4() {
	var sayilar [2][3]int // Satir sayisi //Sutun Sayisi
	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[1][2] = 7
	sayilar[1][1] = 9
	sayilar[1][2] = 5

	fmt.Println(sayilar[1][1])

	fmt.Println()

	//İç İçe For Döngü ile Diziyi yazdırma
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println(sayilar[1][1])
}
