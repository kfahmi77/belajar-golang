package main

import (
	"fmt"
)

func main() {
	var firstName = ""
	lastName := "tanpa tipe data"
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	seventh, eight, ninth := "tujuh", 8, "sembilan"
	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	name := new(string)
	var positiveNumber uint8 = 255
	var decimalNumber float32 = 2.62
	const (
		phi     float32 = 3.14
		istrue          = true
		numeric uint8   = 20
		b
	)
	if b == 100 {
		fmt.Print("b = 100")
	} else {
		fmt.Println("b != 100")
	}
	var poin = 60
	switch poin {
	case 10:
		fmt.Println("lulus dengan nilai sempurna")
	case 8, 9:
		fmt.Println("lulus dengan nilai bagus")
	case 5, 6, 7:
		fmt.Println("lulus")
	default:
		fmt.Println("tidak lulus")

	}

	fmt.Print(b, "\n")
	fmt.Printf("phi: %f \n", phi)
	fmt.Printf("exist? %t \n", false)
	fmt.Printf("bilangan desimal: %.1f\n", decimalNumber)
	fmt.Print(decimalNumber, "\n")
	fmt.Print(positiveNumber, "\n")
	fmt.Print(*name, "aa")
	fmt.Println(name)
	fmt.Print(first, " ", second, third, "\n")
	fmt.Print(seventh, " ", eight, ninth)

}
