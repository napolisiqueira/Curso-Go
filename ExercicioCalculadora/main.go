package main

import "fmt"

func main() {

	var n1 float64
	var n2 float64

	fmt.Println("Digite os dois numeros a serem calculados: ")
	fmt.Scanln(&n1, &n2)

	for {
		choice := menu()
		if choice == 1 {
			calc_sum(n1, n2)
		} else if choice == 2 {
			calc_division(n1, n2)
		} else if choice == 3 {
			calc_minus(n1, n2)
		} else if choice == 4 {
			calc_multiple(n1, n2)
		} else if choice == 5 {
			break
		} else if choice == 0 {
			fmt.Println("Você digitou um valor invalido.")
			continue
		}
	}
}

func menu() int {
	var result int

	teztx := `
	===========================
		    CALCULADORA
	1. SOMA
	2. DIVISÃO
	3. SUBTRAÇÃO
	4. MULTIPLICAÇÃO
	5. SAIR
	==========================`
	fmt.Print(teztx)

	_, err := fmt.Scanln(&result)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return result
}

func calc_sum(number1 float64, number2 float64) {
	result := number1 + number2
	fmt.Println(result)
}

func calc_minus(number1 float64, number2 float64) {
	result := number1 - number2
	fmt.Println(result)
}

func calc_division(number1 float64, number2 float64) {
	result := number1 / number2
	fmt.Println(result)
}

func calc_multiple(number1 float64, number2 float64) {
	result := number1 * number2
	fmt.Println(result)
}
