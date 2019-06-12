package main

import "fmt"

func main() {
	ifTest()
} //en func main
//---------------other functions
func ifTest() {
	var number = 0
	fmt.Println("POR FAVOR INGRESA UN NÚMERO DE 1 A 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("EL NÚMERO ", number, " ES PAR")
	} else {
		fmt.Println("EL NÚMERO ", number, " ES IMPAR")
	}

}
