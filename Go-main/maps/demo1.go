package maps

import "fmt"

//Map Yapisi
func Demo1() {
	//key - value
	//make fonksiyonu ile
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı : ", len(sozluk))

	fmt.Println("Sözluk : ", sozluk)

	delete(sozluk, "book") //Silme
	fmt.Println("Eleman Sayısı : ", len(sozluk))

	fmt.Println("Sözluk : ", sozluk)

	deger := sozluk["table"]
	fmt.Println(deger)

	//true
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu :", varMi)

	//false
	deger, varMii := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu :", varMii)

	sozluk2 := map[string]string{"glass": "bardak", "class": "sinif"}
	fmt.Println(sozluk2)

}
