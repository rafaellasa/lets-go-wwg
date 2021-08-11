package main

import (
	"fmt"
)

func main() {
	//Exercício 1
	fmt.Println("Exercício 1")
	var num1, num2, num3 int = 9, 19, 91
	fmt.Printf("Número 1: %v\nNúmero 2: %v\nNúmero 3: %v\n\n", num1, num2, num3)

	//Exercício 2
	fmt.Println("Exercício 2")
	a := 230
	b := 27
	var sum = a + b
	var sub = a - b
	fmt.Printf("Valor de a: %v\nValor de b: %v\nSoma de a - b: %v\nSubtração de a - b: %v\n\n", a, b, sum, sub)

	//Exercício 3
	fmt.Println("Exercício 3")

	precobanana := 1.25
	precocerveja := 2.98
	precoabacate := 4.59
	precosalgadinho := 7.29

	quantbanana := 2.170
	quantcerveja := 6.0
	quantabacate := 5.650
	quantsalgadinho := 3.0

	fmt.Printf("ITENS\n\n")
	fmt.Printf("Banana: %v quilos * R$ %v = R$ %.2f\n", quantbanana, precobanana, quantbanana*precobanana)
	fmt.Printf("Cerveja: %v unidades * R$ %v = R$ %.2f\n", quantcerveja, precocerveja, quantcerveja*precocerveja)
	fmt.Printf("Abacate: %v quilos * R$ %v = R$ %.2f\n", quantabacate, precoabacate, quantabacate*precoabacate)
	fmt.Printf("Salgadinho: %v unidades * R$ %v = R$ %.2f\n\n", quantsalgadinho, precosalgadinho, quantsalgadinho*precosalgadinho)
	fmt.Printf("TOTAL: R$ %.2f\n\n", quantbanana*precobanana+quantcerveja*precocerveja+quantabacate*precoabacate+quantsalgadinho*precosalgadinho)

	//Exercício 4
	fmt.Println("Exercício 4")
	var nome string = "Rafaella"
	cor := `roxa`
	fmt.Printf("Meu nome é %v e minha cor favorita é %v.\n\n", nome, cor)

	//Exercício 5
	fmt.Println("Exercício 5")

	c := 23 >= 30          // false
	d := 98 < 91           // false
	e := 13 == 13          // true
	f := 15 != 18          // true
	g := 54 > 55 && 9 <= 9 // false

	fmt.Printf("O tipo de a é %T e seu valor é %v\n", c, c)
	fmt.Printf("O tipo de b é %T e seu valor é %v\n", d, d)
	fmt.Printf("O tipo de c é %T e seu valor é %v\n", e, e)
	fmt.Printf("O tipo de d é %T e seu valor é %v\n", f, f)
	fmt.Printf("O tipo de e é %T e seu valor é %v\n\n", g, g)

	//Exercício 6
	fmt.Println("Exercício 6")

	tenhoPreguica := true
	estouComFome := true
	tenhoDinheiro := false

	vouPedirDelivery := tenhoPreguica && estouComFome && tenhoDinheiro //false

	fmt.Printf("Tipo da variável: %T\nResultado: %v\n\n", vouPedirDelivery, vouPedirDelivery)
}
