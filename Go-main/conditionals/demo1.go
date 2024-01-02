package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz!")

	} else if hesap >= cekilmekIstenen {
		fmt.Println("Başarıyla para çekildi")
		hesap -= cekilmekIstenen
	}

	//fmt.Sprintf(hesap)  %f float  formati  degistirme
	fmt.Println("İşlem sonlandı .. Hesaptaki kalan tutar:" + fmt.Sprintf("%f", hesap)) //float
	fmt.Println("İşlem sonlandı .. Hesaptaki kalan tutar:" + fmt.Sprintf("%v", hesap)) //value
	fmt.Printf("İşlem sonlandı .. Hesaptaki kalan tutar:%v", hesap)                    //printf
}
