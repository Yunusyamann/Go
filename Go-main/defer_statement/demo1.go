package defer_statement

import "fmt"

//Defer functions : Bir fonksiyonun en sonunda calisan ve calistigini garanti edilen yapilardir.Loglama yapisi
//Defer en son kutuya girer kutudan ilk cikar mantigi.

func A() {
	fmt.Println("a fonksiyonu calisti.")
}

func C() {
	fmt.Println("c fonksiyonu calisti.")
}

func D() {
	fmt.Println("d fonksiyonu calisti.")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b fonksiyonu calisti.")
     // A()
	 // C()
	 // D()
}
