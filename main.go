package main

import "fmt"

const HelloWorld string = "hola querida(a) %s %s, bienvenido(a) al fascinante mundo de Go! \n" //se puede crear afuera de la funcion
func main() {

	name := getName()
	number := suma(5, 6)
	a, b, c := getVariables()

	//otra forma de declarar variables es ↓ variable := valor, estos 2 puntos + el igual son para declarar una variable sin necesitdad de especificar el tipo de dato
	lastName := getLastName()
	// fmt.Print("INGRESA TU NOMBRE: ")
	// fmt.Scanf("%s", &name) //para obtener un valor de teclado y que se va a gyuardar en la variable name
	fmt.Println("Hola ", name)
	fmt.Printf(HelloWorld, name, lastName) // printf sirve para formatear la cadena de texto con lo q yo necesito, variable y demás cosas
	fmt.Println(number, a, b, c)

} // fin del main

func getName() string {
	var name string
	name = "Sin nombre" //asignacion
	fmt.Print("INGRESA TU NOMBRE: ")
	fmt.Scanf("%s", &name)
	return name
}

func getLastName() string {
	var lastName string
	lastName = "Sin apellido"
	fmt.Print("INGRESA TU APELLIDO: ")
	fmt.Scanf("%s", &lastName)
	return lastName
}

func getVariables() (int, int, int) {
	return 1, 2, 3
} //func getVariables

func suma(a int, b int) int {
	return a + b
} //func suma
