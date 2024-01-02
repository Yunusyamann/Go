package string_functions

//alias
import (
	"fmt"
	s "strings"
)

// Case Sensitive (Buyuk-Kucuk Harf Duyarli)
// ascii (karakterler)
func Demo1() {
	isim := "Yunus"
	fmt.Println(s.Count(isim, "n"))    //Count : n harfi kac defa var.
	fmt.Println(s.Contains(isim, "n")) //Contains : 1-0 harf var mi yok mu
	fmt.Println(s.Index(isim, "n"))    //Index: Kacinci index. 4
	fmt.Println(s.Index(isim, "w"))    // -1 bulamadiginda dondurur.

	sonuc := s.Index(isim, "a")
	if sonuc != -1 {
		fmt.Println("a harfi var.")
	} else {
		fmt.Println("a harfi yok.")
	}

	fmt.Println(s.ToLower(isim)) // ToLower:Kucuk harfe donusturme
	fmt.Println(s.ToUpper(isim)) // ToUpper:Buyuk harfe donusturme

}
