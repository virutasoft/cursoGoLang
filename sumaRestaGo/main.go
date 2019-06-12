package main

import "fmt"

func main() {
	var option int
	fmt.Print("---Hola, bienvenido a la calculadora básica en GoLang---- \n")

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
func suma(a float64, b float64) float64 {
	return a + b
} // end func suma

func resta(a float64, b float64) float64 {
	return a - b
} //end func resta
func multiplicar(a float64, b float64) float64 {
	return a * b
} // end multiplicar
func division(a float64, b float64) float64 {
	return a / b
}
func getNumbers() (float64, float64) {
	var (
		a float64
		b float64
	)
	fmt.Print("Ingresa la cifra 1: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Ingresa la cifra 2: ")
	fmt.Scanf("%f", &b)
	return a, b
} //end getNumbers
