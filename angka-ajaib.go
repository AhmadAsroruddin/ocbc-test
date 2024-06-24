package main

import (
	"fmt"
	"strconv"
)

func angkaAjaib(n string) int {
	constantaKaprekars := 0
	for n != "6174" {
		ascending, _ := strconv.Atoi(sorting(n, true))
		descending, _ := strconv.Atoi(sorting(n, false))
		n = strconv.Itoa(descending - ascending)
		if len(n) < 4 {
			n = fmt.Sprintf("%04s", n)
		}
		constantaKaprekars++
	}

	return constantaKaprekars
}

func sorting(n string, isAscending bool) string {
	numArr := stringToArr(n)
	numArrLen := len(numArr)

	for i := 0; i < numArrLen-1; i++ {
		for j := 0; j < numArrLen-1-i; j++ {
			if isAscending {
				if numArr[j] > numArr[j+1] {
					numArr[j], numArr[j+1] = numArr[j+1], numArr[j]
				}
			} else { //desc
				if numArr[j] < numArr[j+1] {
					numArr[j], numArr[j+1] = numArr[j+1], numArr[j]
				}
			}
		}
	}

	return arrToString(numArr)
}

func stringToArr(n string) []string {
	digits := make([]string, len(n))

	for i, value := range n {
		digits[i] = string(value)
	}
	return digits
}

func arrToString(arr []string) string {
	var result string
	for _, value := range arr {
		result += value
	}

	result = fmt.Sprintf("%04s", result)
	return result
}
