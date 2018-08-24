package utils

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

func IsPrime(num int) bool {
	maxVal := int(math.Sqrt(float64(num))) + 1

	for x := 2; x < maxVal; x++ {
		if num%x == 0 {
			return false
		}
	}

	return true
}

func Fib(limit int) []int {
	first := 1
	second := 2

	var numbers []int

	numbers = append(numbers, first)
	numbers = append(numbers, second)

	for {
		third := first + second

		if third > limit {
			break
		} else {
			numbers = append(numbers, third)

			first = second
			second = third
		}
	}

	return numbers
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func IsPalindrome(num int) bool {
	var digits []int
	temp := num

	for temp > 0 {
		digit := temp % 10
		digits = append(digits, digit)
		temp = int(temp / 10)
	}

	digitCount := len(digits)

	for i := 0; i < digitCount; i++ {
		reversei := digitCount - i - 1

		if i < reversei && digits[i] != digits[reversei] {
			return false
		}
	}

	return true
}

func IsNaturalNumber(num float64) bool {
	return false
}

func BoardColumn(board [][]float64, columnIndex int) []float64 {
	column := make([]float64, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return column
}

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 3, 64)
}

func SliceFtoa(sa []float64) []string {
	si := make([]string, 0, len(sa))

	for _, a := range sa {
		i := FloatToString(a)
		si = append(si, i)
	}

	return si
}

func MakeCSV(points [][]float64) {
	file, err := os.Create("result.csv")
	CheckError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range points {
		err := writer.Write(SliceFtoa(value))
		CheckError("Cannot write to file", err)
	}
}

func MultiplySlice(sli []int) int {
	product := 1

	for _, num := range sli {
		product *= num
	}

	return product
}

func FindDivCount(num int) int {
	divCount := 2

	for div := 2; div <= int(math.Sqrt(float64(num))); div++ {
		if num%div == 0 {
			divCount += 2
		}
	}

	return divCount
}

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func GetDigitFromString(str string, position int) int {
	digit, _ := strconv.Atoi(string(str[position]))
	return digit
}

func PadStringWithZero(value string, zeroCount int) string {
	result := value
	for x := 0; x < zeroCount; x++ {
		result = "0" + result
	}

	return result
}

func AddStringsAsNumbers(first string, second string) string {
	result := ""

	if len(second) > len(first) {
		temp := first
		first = second
		second = temp
	}

	second = PadStringWithZero(second, len(first)-len(second))

	carry := 0

	for x := len(first) - 1; x >= 0; x-- {
		firstDigit := GetDigitFromString(first, x)
		secondDigit := GetDigitFromString(second, x)

		currentSum := firstDigit + secondDigit + carry

		if currentSum >= 10 {
			carry = int(currentSum / 10)
			currentSum = currentSum % 10
		} else {
			carry = 0
		}

		result += strconv.Itoa(currentSum)
	}

	if carry != 0 {
		result += strconv.Itoa(carry)
	}

	return Reverse(result)
}
