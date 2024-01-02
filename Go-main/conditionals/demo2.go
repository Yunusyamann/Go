package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz!")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Hesaptaki Paranız Hazır!")
		fmt.Println("Dikkat Hesabınızda para kalmadı.!")
	} else {
		fmt.Println("Hesaptaki Paranız Hazırlanıyor!")
	}
}
