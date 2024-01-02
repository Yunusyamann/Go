package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "Istanbul"
	sehirler[2] = "Hatay"
	sehirler[3] = "Eskişehir"
	sehirler[4] = "Izmir"
	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}

}
