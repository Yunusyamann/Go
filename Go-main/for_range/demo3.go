package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k := range sozluk {
		fmt.Println(k)
	}

	for k, v := range sozluk {
		fmt.Println(k) //key
		fmt.Println(v) //value
	}
}
