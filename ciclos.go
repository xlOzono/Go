package main

import "fmt"

func main() {
    fmt.Print("Ingrese un número límite: ")
    var limite int
    fmt.Scan(&limite)

    fmt.Printf("Números perfectos hasta %d:\n", limite)
    for i := 1; i <= limite; i++ {
        if esPerfecto(i) {
            fmt.Println(i)
        }
    }
}

func esPerfecto(num int) bool {
    sum := 0
    for i := 1; i < num; i++ {
        if num%i == 0 {
            sum += i
        }
    }
    return sum == num
}
