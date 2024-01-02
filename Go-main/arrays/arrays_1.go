package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int //[0 0 0 0 0]
	fmt.Println(sayilar)

	fmt.Println()

	sayilar[2] = 50
	fmt.Println(sayilar[2])
}
