package loops

import "fmt"

func Demo1() {
	var metin string = "Hello World"
	i := 1 //index
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}
	fmt.Println("Bitti")

	//infinite loop :Sonsuz Döngü
	/*for i <= 5 {
		fmt.Println(metin)
	}
	*/
}
