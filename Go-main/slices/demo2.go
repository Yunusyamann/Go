package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler)) // MAKE ile yeni dizi olusturuldu.
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler) //slice başka slice atama
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "Hatay")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	//Slice Okuma
	fmt.Println(sehirler[1:3]) //Indis 1 den 3 e kadar dahil degil
	fmt.Println(sehirler[1:5])

	fmt.Println(sehirler[1:]) //Indis 1 den tumu

	fmt.Println(sehirler[:2]) //indis 0 dan 2 ye kadar

}
