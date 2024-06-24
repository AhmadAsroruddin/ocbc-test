package main

import "fmt"

func perfectTwo(arr []int, target int) string {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return fmt.Sprintf("[%d,%d]", i, j)
			}
		}
	}

	return "<No Way>"
}
