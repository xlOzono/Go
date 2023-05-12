package main

import "fmt"

func main() {

	fmt.Print("Ingrese el numero de personas: ")
	var personas int
	fmt.Scan(&personas)

	fmt.Print("Ingrese el numero de galletas: ")
	var galletas int
	fmt.Scan(&galletas)

	galletasPorAmigo := galletas / personas
	galletasSobrantes := galletas % personas

	fmt.Printf("Se deben dar %d galletas a cada amigo\n", galletasPorAmigo)
	fmt.Printf("Sobran %d gallentas \n", galletasSobrantes)
}
