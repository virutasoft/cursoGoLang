package main

import "fmt"

const palabra string = "Rogletronicon"

func main() {

	fmt.Print("zorra \n")
	stringUTF8 := getUnicode()
	fmt.Println("CADENA CON UTF8: ", stringUTF8)
	// para que al imprimir los indices de una palabra o frase, me devuelva la letra en la posición deseada, debo añadile string a la impresión, sino, traerá el código ascii de el caracter
	fmt.Println(string("HOLA"[0]))
	//para saber la cantidad de caracteres o cifras en una frase, usamos "len"
	fmt.Println("LA CANTIDAD DE CARACTERRES QUE TIENE LA PALABRA: ", palabra, " es ", len(palabra))
}

func getUnicode() string {
	return "人物"
}
