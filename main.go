package main

import (
	"fmt"
	"time"
	// "github.com/google/uuid"
)

func hola() {
	time.Sleep(1 * time.Second)
	fmt.Println("hola mundo")
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	go hola()
	// fmt.Println(uuid.NewString())

	encontrados := 0
	numero := 2

	fmt.Println("Los primeros 100,000 números primos son:")

	for encontrados < 100 {
		if esPrimo(numero) {
			fmt.Println(numero)
			encontrados++
		}
		numero++
	}
}
