package main

import "fmt"

func main() {
	var n int
	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&n)

	fmt.Println("Los numeros perfectos hasta",n,"son: ")

	for i := 1; i <= n; i++{
		if esPerfecto(i, i - 1){
			fmt.Println(i)
		}
	} 

}

func esPerfecto(num, divisorSum int) bool{
	if divisorSum == 0{
		return num == 0
	}

	if num % divisorSum == 0{
		return esPerfecto(num - divisorSum, divisorSum - 1)
	}

	return esPerfecto(num, divisorSum - 1)
}