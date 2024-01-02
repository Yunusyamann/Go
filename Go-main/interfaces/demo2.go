package interfaces

import "fmt"

//Gercek Hayat Ornegi Banka Kredi ürünler

type Mortgage struct {
	creditPaymentTotal float64
	addres             string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	CarInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

//Polymorfizm
func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, addres: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, addres: "Istanbul"}
	credits3 := Car{rate: 15, creditPaymentTotal: 300000, CarInfo: "Mercedes"}

	credits := []CreditCalculator{credit1, credit2, credits3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Toplam Ödeme :", total)

}
