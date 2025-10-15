package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	if 8%4 == 0 {
		fmt.Println("8 e par")
	} else {
		fmt.Println("8 e impar")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("um dos dois e par")
	}

	if num := 9; num < 0 {
		fmt.Println("numero e menor que 0")
	} else if num < 10 {
		fmt.Println("numero e menor que 10")
	} else {
		fmt.Println("numero e maior que 10")
	}
}
