package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo2() {
	f, err := os.Open("demo11.txt") //nil Hata yok.
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok { //dosyanın yoluyla ilgili bir hata olması.
			fmt.Println("Dosya Bulunamadi : ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadi")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
