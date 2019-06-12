package main

import "fmt"

func main() {
	getArray()
	getSlice()
}

func getArray() {
	//solo se puede imprimir los elementos que tiene el array, nada mas
	var array1 [4]string
	//segunda forma de declarar arrays
	array2 := [4]int{1, 2, 3, 4}
	array1[0] = "camilo "
	array1[1] = "es "
	array1[2] = "una "
	array1[3] = "loca "
	fmt.Println(array1, array2)
}

//SLICE: no tienen una cantidad de elementos fija, es dinámico
func getSlice() {
	var slice1 []string
	//sintaxis de un slice(array dinamico)
	/*nombreArray = append(nombreArray, "xxx",111,"aaa",...*/
	slice1 = append(slice1, "el", "calvo", "es", "un", "maricon")
	fmt.Println(slice1)

}


