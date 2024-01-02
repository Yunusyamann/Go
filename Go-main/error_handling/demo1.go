package error_handling

import (
	"fmt"
	"os"
)

// Hata Yakalama: Programcinin hatayi kontrol edemedigi fakat hataya gore aksiyon alan yapilardir.
// Defensive Programming
func Demo1() {
	f, err := os.Open("demo1.txt") //demo1.txt dosyasini acma islemi
	// f dosya bulunursa  Nil 0 döndürür

	if err != nil {
		fmt.Println("Dosya Bulunamadi")
	} else {
		fmt.Println(f.Name())
		//dosyanın ismini yaz

	}

}
