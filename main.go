package main

import (
	"fmt"
	"strings"

	"github.com/tecnologer/romantoarabic/num"
)

var err error

func main() {
	var command int
	fmt.Println("Opciones: ")
	fmt.Println("1.- Romano a Arabigo")
	fmt.Println("2.- Arabigo a Romano")

	for err != nil || command < 1 || command > 2 {
		fmt.Print("Escribe una opcion correcta: ")
		_, err = fmt.Scan(&command)
	}

	switch command {
	case 1:
		var romanInput string
		var isCorrect bool
		fmt.Print("Escribe el numero romano correcto: ")
		for err != nil || romanInput == "" || !isCorrect {
			_, err = fmt.Scan(&romanInput)
			isCorrect = num.IsRoman(romanInput)
			if !isCorrect {
				fmt.Printf("\"%s\" no es un numero romano, intenta de nuevo: ", romanInput)
			}
		}

		fmt.Printf("%s es igual a %d\n", strings.ToUpper(romanInput), num.ParseRoman(romanInput))
	case 2:
		var arabicInput int
		fmt.Print("Escribe el numero arabigo correcto (mayor a 0 y menor a 4000): ")
		for err != nil || arabicInput < 1 || arabicInput > 3999 {
			_, err = fmt.Scan(&arabicInput)
		}

		fmt.Printf("%d es igual a %s\n", arabicInput, num.ParseArabic(arabicInput))
	}
	// var arabicInput = 3999
	// fmt.Printf("%d es igual a %s", arabicInput, ParseArabic(arabicInput))
}
