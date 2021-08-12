package main

import (
	"fmt"
	"time"
)

func main() {
	//Exercício 1

	//O programa não compila porque o tipo int8 representa valores no intervalo -128 a 127,
	//e a variável recebeu o valor 150, maior que o limite máximo. Para corrigir, o tipo
	//int8 da variável deve ser alterado para um que comporte o valor 150, por exemplo, int16.
	var quilometros int16
	quilometros = 150
	fmt.Printf("Exercício 1\n")
	fmt.Println(quilometros)

	//Exercício 2
	var ano int
	fmt.Printf("\nExercício 2")
	fmt.Printf("\nDigite o ano de seu nascimento: ")
	fmt.Scan(&ano)
	fmt.Printf("Sua idade atual é %v.\n\n", time.Now().Year()-ano)
}
