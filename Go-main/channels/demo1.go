package channels

import (
	"fmt"
	"time"
)

// Channels :Asenkron çalışmasını ve yaşam döngüsünü kontrol eder
func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
		time.Sleep(1 * time.Second)
		fmt.Println("Çift sayı toplama çalışıyor")
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
		time.Sleep(1 * time.Second)
		fmt.Println("Tek sayı toplama  çalışıyor")
	}
	tekSayiCn <- toplam
}
