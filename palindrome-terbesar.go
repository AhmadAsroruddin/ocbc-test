package main

import (
	"strconv"
)

func isPalindrome(num string) bool {
	n := len(num)
	for i := 0; i < n/2; i++ {
		if num[i] != num[n-1-i] {
			return false
		}
	}

	return true
}

func palindromeTerbesar(n int) string {
	if n <= 4 {
		maxPal := 0
		upperLimit := 1
		lowerLimit := 1

		for i := 0; i < n; i++ {
			upperLimit *= 10
			if i != 0 {
				lowerLimit *= 10
			}
		}

		for i := upperLimit - 1; i >= lowerLimit; i-- {
			for j := i; j >= lowerLimit; j-- {
				num := i * j
				if isPalindrome(strconv.Itoa(num)) && num > maxPal {
					maxPal = num
				}
			}
		}
		return strconv.Itoa(maxPal)
	} else {
		return "Input must be lower than 5"
	}

}
