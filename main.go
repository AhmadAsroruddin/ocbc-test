package main

import "fmt"

func main() {
	//HOT DOG

	//UNCOMMENT THIS TO RUN HOT DOT
	//wordDict := []string{"hot", "dot", "dog", "lot", "log"}
	//
	//fmt.Println(hotDog(wordDict, "hit", "dig"))

	//PROGRAMMER

	//UNCOMMENT THIS TO RUN PROGRAMMER
	wordDict := []string{"pro", "gram", "merit", "program", "it", "programmer"}
	output := wordBreak("programmerit", wordDict)
	for _, way := range output {
		fmt.Println(way)
	}
}
