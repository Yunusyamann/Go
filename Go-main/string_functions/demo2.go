package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isimSoyisim := "Yunus Yaman"
	fmt.Println(s.HasPrefix(isimSoyisim, "Yun")) //HasPrefix:On Ek
	fmt.Println(s.HasSuffix(isimSoyisim, "ama")) //HasSuffix:Son Ek

	fmt.Println(s.Index(isimSoyisim, "Yu")) //Kacinci Index

	harfler := []string{"Y", "u", "n", "u", "s"}
	fmt.Println(s.Join(harfler, "*")) //Join ile Birlestirme islemi

	harfler2 := []string{"Y", "u", "n", "u", "s"}
	sonuc := (s.Join(harfler2, "-")) //Join
	fmt.Println(sonuc)

	harfler3 := []string{"Y", "u", "n", "u", "s"}
	sonuc1 := (s.Join(harfler3, ""))            //Join
	fmt.Println(s.Replace(sonuc1, "", "+", -1)) //Replace:Karakter degistirme  -1 Tumunu degistir
	fmt.Println(s.Replace(sonuc1, "", "+", 3))  //En fazla 3 karakter degistir
	fmt.Println(s.Replace(sonuc1, "", "+", 1))

	//Split: Join Tam Tersi Belli formata gore ayirma parcalama yontemi
	fmt.Println(s.Split(sonuc1, "-"))

	fmt.Println(s.Repeat(sonuc1, 5)) //Repeat:Tekrarlama Kopyasini yazdirma

}
