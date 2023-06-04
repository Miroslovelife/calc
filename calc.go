package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var num1 int
var num2 int

var oper string
var result int
var rimArray = []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func main() {

	fmt.Print("Введите выражение: ")

	//Считование строки
	var inputStr string
	fmt.Fscan(os.Stdin, &inputStr)
	if len(inputStr) < 3 {
		panic("Строка не является математической операцией.")
	}

	//Преобразование строки(без пробелов)
	//inputStr = strings.Replace(inputStr, " ", "", 2)

	inputStrWithoutOper := regexp.MustCompile("[+,\\-,*,/]{1}").Split(inputStr, -1)
	if len(inputStrWithoutOper) > 2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	oper = detectOper(inputStr)

	if len(oper) == 0 {
		fmt.Println("Строка не является математической операцией.")
	} else if len(oper) > 1 {
		fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	if isRoman(inputStrWithoutOper[0]) && isRoman(inputStrWithoutOper[1]) {

		result = calc(convertToArab(inputStrWithoutOper[0], inputStrWithoutOper[1]))
		if result < 0 {
			fmt.Println("Римской системе нет отрицательных чисел.")
		} else if result == 0 {
			fmt.Println("0 не может быть результатом работы калькулятора с римскими числами ")
		} else {
			fmt.Println(convertToRoman(result))
		}

	} else if !isRoman(inputStrWithoutOper[0]) && !isRoman(inputStrWithoutOper[1]) {
		numToInt, err := strconv.Atoi(inputStrWithoutOper[0])
		if err != nil {
			panic("")
		}
		num1 = numToInt

		numToInt1, err := strconv.Atoi(inputStrWithoutOper[1])
		if err != nil {
			panic("")
		}
		num2 = numToInt1

		fmt.Println(calc(num1, num2))
	} else {
		fmt.Println("Используются одновременно разные системы счисления.")
	}
}

func detectOper(inputStr string) string {

	for liters := 0; liters < len(inputStr); liters++ {
		if string(inputStr[liters]) == "+" {
			oper = "+"
		} else if string(inputStr[liters]) == "-" {
			oper = "-"
		} else if string(inputStr[liters]) == "*" {
			oper = "*"
		} else if string(inputStr[liters]) == "/" {
			oper = "/"
		}

	}
	return oper
}

func isRoman(val string) bool {

	var clown bool
	for i := 0; i < len(rimArray); i++ {

		if val == rimArray[i] {
			clown = true
			break
		} else {
			clown = false
		}

	}
	return clown
}

func convertToArab(num1, num2 string) (int, int) {
	var resultArabConvert int
	var resultArabConvert_ int
	for i := 0; i < len(rimArray); i++ {
		if num1 == rimArray[i] {
			resultArabConvert = i
		}

	}
	for i := 0; i < len(rimArray); i++ {
		if num2 == rimArray[i] {
			resultArabConvert_ = i
		}

	}
	return resultArabConvert, resultArabConvert_
}

func convertToRoman(ArabNum int) string {
	var resultRomanConvert = ""
	resultRomanConvert = rimArray[ArabNum]

	return resultRomanConvert
}

func calc(num1, num2 int) int {
	if oper == "+" {
		result = num1 + num2
	} else if oper == "-" {
		result = num1 - num2
	} else if oper == "*" {
		result = num1 * num2
	} else if oper == "/" {
		result = num1 / num2
	}
	return result
}
