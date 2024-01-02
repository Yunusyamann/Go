package main //Paket tanımlama
import (
	"fmt"
	"golesson/project"
)

//"golesson/interfaces"

func main() {
	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo1()
	//conditionals.Demo2()
	// conditionals.Demo3()
	//conditionals.Demo4()
	//loops.Demo1()
	// loops.Demo2()
	//loops.Demo3()
	//loops.Demo4()
	//loops.Demo5()
	//loops.Demo6()
	//loops.Demo7()
	//loops.Demo8()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//arrays.Demo4()
	//slices.Demo1()
	//slices.Demo2()

	//functions.SelamVer("Ozcan")
	//functions.Topla(2, 6)

	/*
		var sonuc = functions.Toplam(3, 8)
		fmt.Println(sonuc * 10)
	*/

	/*
		sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

		fmt.Println("Toplam :", sonuc1)
		fmt.Println("Cikarma :", sonuc2)
		fmt.Println("Carpim :", sonuc3)
		//fmt.Println("Bolum :", sonuc4)
	*/

	/*
		var sonuc = functions.ToplaVariadic(12, 4, 7, 9, 87, 65)
		fmt.Println(sonuc)

		fmt.Println(functions.ToplaVariadic(12, 4, 7))
		fmt.Println(functions.ToplaVariadic())

		sayilar := []int{4, 6, 7, 30, 11}
		fmt.Println(functions.ToplaVariadic(sayilar...))
	*/

	//maps.Demo1()

	//for_range.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()

	/*
		sayi := 20
		pointers.Demo1(&sayi)
		fmt.Println("Maindeki sayı :", sayi)
	*/

	/*
		sayilar := []int{1, 2, 3}
		pointers.Demo2(sayilar)
		fmt.Println("Maindeki sayı :", sayilar[0])
	*/

	//structs.Demo1()
	//structs.Demo2()

	/*
		goroutines.CiftSayilar()
		goroutines.TekSayilar()
	*/

	/*
		go goroutines.CiftSayilar()
		go goroutines.TekSayilar()
		time.Sleep(5 * time.Second) //5 Saniye
		fmt.Println("Main Bitti")
	*/

	/*
		ciftSayiCn := make(chan int)
		tekSayiCn := make(chan int)
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

		carpim := ciftSayiToplam * tekSayiToplam
		fmt.Println("Çarpım :", carpim)
	*/

	//interfaces.Demo1()
	//interfaces.Demo2()

	// defer_statement.B()
	//defer_statement.Test()
	//defer_statement.Demo3()

	//error_handling.Demo1()
	//error_handling.Demo2()

	//interfaces.Demo3()

	//error_handling.Demo3()
	//fmt.Println(error_handling.TahminEt2(102))

	//string_functions.Demo1()
	//string_functions.Demo2()

	//restful.Demo1()
	// restful.Demo2()

	//project.GetAllProducts()
	//project.AddProduct()

	//products, _ := project.GetAllProducts()
	//fmt.Println(products)

	//Verileri tek tek yazdirma
	/*
		for i := 0; i < len(products); i++ {
			fmt.Println(products[i].ProductName)
		}
	*/

	/*
		project.AddProduct()
		products, _ := project.GetAllProducts()
		for i := 0; i < len(products); i++ {
			fmt.Println(products[i].ProductName)
		}
	*/

	product, _ := project.GetAllProducts()
	fmt.Println(product)

	project.AddProduct()
	products, _ := project.GetAllProducts()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}
