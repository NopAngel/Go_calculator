package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Ingresa el primer número: ")
	fmt.Scan(&num1)

	fmt.Print("Ingresa el segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Selecciona la operación (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: No se puede dividir por cero.")
		}
	default:
		fmt.Println("Operador no válido. Por favor, selecciona +, -, *, o /.")
	}
}