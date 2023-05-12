package main

import "fmt"

func main() {
	fmt.Print("Ingrese numerador: ")
	var numerador int
	fmt.Scan(&numerador)

	fmt.Print("Ingrese denominador: ")
	var denominador int
	fmt.Scan(&denominador)

	if numerador < denominador{
		fmt.Println("Es una fraccion propia")
	}

	else if numerador % denominador == 0{
		fmt.Println("Es una fraccion propia, su resultado es: ", numerador / denominador)
	}

	else{
		entero := numerador / denominador
		resto := numerador % denominador
		fmt.Printf("Es una fraccion impropia y su forma mixta es: %d %d/%d\n",entero,resto,denominador)
	}
}