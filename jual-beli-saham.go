package main

import (
	"math"
	"strconv"
)

func jualBeliSaham(price []int) string {
	lowest := math.MaxInt
	maxProfit := 0
	downCount := 0

	for _, value := range price {
		if value < lowest {
			lowest = value
			downCount++
		}

		profit := value - lowest
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	if downCount == len(price) {
		return "Harga Selalu Turun"
	} else {
		return strconv.Itoa(maxProfit)
	}
}
