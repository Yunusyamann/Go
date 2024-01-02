package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Urun Kaydedildi : ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Log islemi uygulandi.")
}

func Demo3() {
	p := product{productName: "MacBook", unitPrice: 25000}
	p = product{productName: "Iphone14", unitPrice: 40000}

	fmt.Println("Islem Basarili")
	fmt.Println(p.productName)
	defer p.save() //Son Hali print etmek icin
}
