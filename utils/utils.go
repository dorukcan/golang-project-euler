package utils

import "math"

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
