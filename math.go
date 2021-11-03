package main

import "fmt"

func main() {
	resultSum := Sum(10, 10)

	fmt.Println("The result of sum is:", resultSum)
}

func Sum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}
