package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	const titleart = `   	 ____            _         _____      _            _       _             
	|  _ \          (_)       / ____|    | |          | |     | |            
	| |_) | __ _ ___ _  ___  | |     __ _| | ___ _   _| | __ _| |_ ___  _ __ 
	|  _ < /  ' / __| |/ __| | |    / _' | |/ __| | | | |/ _' | __/ _ \| '__|
	| |_) | (_| \__ \ | (__  | |___| (_| | | (__| |_| | | (_| | || (_) | |   
	|____/ \__,_|___/_|\___|  \_____\__,_|_|\___|\__,_|_|\__,_|\__\___/|_|   
																			 `

	fmt.Println("\n\u001b[33m", titleart, "\u001b[0m")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\u001b[32m")
		fmt.Println("[1] Sumar")
		fmt.Println("[2] Restar")
		fmt.Println("[3] Multiplicar")
		fmt.Println("[4] Dividir")
		fmt.Println("[5] Salir")
		fmt.Print("\u001b[0m")

		fmt.Print("\nDigita una de las opciones: ")
		scanner.Scan()

		opcion, opcerr := strconv.Atoi(scanner.Text())

		if opcerr != nil || opcion < 1 || opcion > 5 {
			fmt.Println("\n\n\u001b[31mElija una de las 5 opciones que se le muestran. Evite hacer lo contrario de lo que se le pide!!! >:(\u001b[0m\n\n")
			time.Sleep(1500 * time.Millisecond)
		} else if opcion != 5 {

			fmt.Println("\nDigita los valores hermano :)\n")
			fmt.Print("Valor 1: ")
			scanner.Scan()
			num1, num1err := strconv.Atoi(scanner.Text())
			fmt.Print("Valor 2: ")
			scanner.Scan()
			num2, num2err := strconv.Atoi(scanner.Text())

			if num1err != nil || num2err != nil {
				fmt.Println("\n\u001b[31mVete a la burger, no pienso ayudarte si te comportas como un imbecil UnU\u001b[0m")
			} else {

				fmt.Print("\nLa respuesta es: ")

				switch opcion {
				case 1:
					fmt.Println(suma(num1, num2))
					break
				case 2:
					fmt.Println(resta(num1, num2))
					break
				case 3:
					fmt.Println(multiplicacion(num1, num2))
					break
				case 4:
					fmt.Println(division(num1, num2))
					break
				default:
					fmt.Println("\n\u001b[33mNo se que, pero algo hiciste ma compa :/\u001b[0m")
					break
				}

				time.Sleep(1500 * time.Millisecond)
			}

		}

		if opcion == 5 {
			break
		}
	}

}

func suma(n1 int, n2 int) int {
	return n1 + n2
}

func resta(n1 int, n2 int) int {
	return n1 - n2
}

func multiplicacion(n1 int, n2 int) int {
	return n1 * n2
}

func division(n1 int, n2 int) int {
	return n1 / n2
}
