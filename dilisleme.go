package main

import (
	"fmt"

	Tahmin "github.com/endeveit/guesslanguage"
)

func main() {
	fmt.Print("Dil'i tespit edilecek kelime veya cümleyi giriniz : ")
	fmt.Scanln("&girdi")
	fmt.Println(Tahmin.Guess(girdi))
}
