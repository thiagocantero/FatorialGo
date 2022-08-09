package main

import "fmt"

func main() {

	var num_fatorial, fatorial int
	fatorial = 1

	fmt.Print("Digite o número para calcular seu fatorial = ")
	fmt.Scanln(&num_fatorial)

	for i := 1; i <= num_fatorial; i++ {
		fatorial = fatorial * i
	}
	fmt.Println("A Fatorial de ", num_fatorial, " é = ", fatorial)
}
