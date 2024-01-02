package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("Cift sayi calisti")
		return "Çift Sayidir."
	}
	if sayi%2 != 0 {
		return "Tek Sayidir."
	}
	DeferFunc()
	return "Belli degil"
}
func Test() {
	sonuc := Demo2(11)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("İslem Bitti")
}
