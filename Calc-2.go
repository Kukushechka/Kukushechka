package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	myscanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите числа для решения римскими или арабскими цифрами: ")
	myscanner.Scan()
	input := myscanner.Text()
	var variables [3]string
	variables = [3]string(strings.Split(input, " "))
	if len(variables) > 3 {
		panic("Слишком много")
	}

	var number1, number2 int
	result := 0
	if isItRoman(variables[0]) && isItRoman(variables[2]) {
		number1 = romanToArab(variables[0])
		number2 = romanToArab(variables[2])
	} else {
		if isItArab(variables[0]) && isItArab(variables[2]) {
			x, _ := strconv.Atoi(variables[0])
			y, _ := strconv.Atoi(variables[2])
			number1 = x
			number2 = y

		} else {
			panic("BREAK")
		}
	}
	switch variables[1] {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
		if result <= 0 {
			panic("PANIK")
		}
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		panic("BEAK")
	}

	if isItRoman(variables[0]) && isItRoman(variables[2]) {
		resultArab := arabInRoman(result)
		fmt.Print(resultArab)
	} else {
		fmt.Println(result)
	}
}

func isItRoman(number string) bool {

	romanNumbersMap := make(map[int]string)

	rumNumbLow := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	rumNumbBIG := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}

	for i := 0; i < len(rumNumbLow); i++ {
		romanNumbersMap[i] = rumNumbLow[i]
	}
	l := 9
	for i := 1; i <= l; i++ {
		for j := 1; j <= l; j++ {
			romanNumbersMap[(i)*10+j] = rumNumbBIG[i-1] + rumNumbLow[j] // что то здесь наворотил
		}
	}
	for i := 0; i < 10; i++ {
		for j := 10; j < 100; j++ {
			romanNumbersMap[(i+1)*10] = rumNumbBIG[i]
		}
	}
	// создал мапу с римскими числами относительно числовых
	for i := 0; i < 3; i++ {
		for k := 0; k < len(romanNumbersMap); k++ {
			if number == romanNumbersMap[k] {

				return true
			}
		}
	}
	return false
}

func arabInRoman(result int) string {
	romanNumbersMap := make(map[int]string)

	rumNumbLow := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	rumNumbBIG := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}

	for i := 0; i < len(rumNumbLow); i++ {
		romanNumbersMap[i] = rumNumbLow[i]
	}
	l := 9
	for i := 1; i <= l; i++ {
		for j := 1; j <= l; j++ {
			romanNumbersMap[(i)*10+j] = rumNumbBIG[i-1] + rumNumbLow[j] // что то здесь наворотил
		}
	}
	for i := 0; i < 10; i++ {
		for j := 10; j < 100; j++ {
			romanNumbersMap[(i+1)*10] = rumNumbBIG[i]
		}
	}
	resultRoman := romanNumbersMap[result-1]
	return resultRoman
}

func romanToArab(num1 string) int {
	romanNumbersMap := make(map[int]string)

	rumNumbLow := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	rumNumbBIG := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}

	for i := 0; i < len(rumNumbLow); i++ {
		romanNumbersMap[i] = rumNumbLow[i]
	}
	l := 9
	for i := 1; i <= l; i++ {
		for j := 1; j <= l; j++ {
			romanNumbersMap[(i)*10+j] = rumNumbBIG[i-1] + rumNumbLow[j] // что то здесь наворотил
		}
	}
	for i := 0; i < 10; i++ {
		for j := 10; j < 100; j++ {
			romanNumbersMap[(i+1)*10] = rumNumbBIG[i]
		}
	}

	number := 0

	for i := 0; i < 3; i++ {
		for k := 0; k < len(romanNumbersMap); k++ {
			if num1 == romanNumbersMap[k] {
				number = k + 1
				break
			}
		}
	}
	return number
}

func isItArab(numb string) bool {
	var arabNumb = make([]int, 10)
	for i := 0; i < len(arabNumb); i++ {
		arabNumb[i] = i + 1
	}

	number1, _ := strconv.Atoi(numb)

	for i := 0; i < len(arabNumb); i++ {

		if number1 == arabNumb[i] {
			return true
		}
	}
	return false
}
