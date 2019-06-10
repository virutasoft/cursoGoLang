package main

import "fmt"

func main() {
	var option int
	fmt.Print("---Hola, bienvenido a la calculadora de sumas y restas---- \n")

	fmt.Print("Si deseas sumar, escribe: 1, 2 si deseas restar, 3 si deseas multiplicar o 4 si deseas dividir \n")
	fmt.Scanf("%d", &option)
	if option == 1 {
		suma := suma(getNumbers())
		fmt.Println("El resultado de la suma es: ", suma)
	} else if option == 2 {
		resta := resta(getNumbers())
		fmt.Println("El resultado de la resta es: ", resta)
	} else if option == 3 {
		multiplicar := multiplicar(getNumbers())
		fmt.Println("El resultado de la multiplicacion es: ", multiplicar)
	} else if option == 4 {
		division := division(getNumbers())
		fmt.Println("El resultado de la division es: ", division)
	} else {
		fmt.Println("no has ingresado la opcion correcta")
	}
} //end main
func suma(a int, b int) int {
	return a + b
} // end func suma

func resta(a int, b int) int {
	return a - b
} //end func resta
func multiplicar(a int, b int) int {
	return a * b
} // end multiplicar
func division(a int, b int) int {
	return a / b
}
func getNumbers() (int, int) {
	var (
		a int
		b int
	)
	fmt.Print("Ingresa la cifra 1: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Ingresa la cifra 2: ")
	fmt.Scanf("%d", &b)
	return a, b
} //end getNumbers
