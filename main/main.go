package main

import (
	"fmt"
	"os"

	"reloaded"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Need the samples.txt results.text")
		return
	}

	textFile := args[0]
	outputFile := args[1]

	textContent, err := reloaded.ReadTextFile(textFile)
	if err != nil {
		fmt.Println("Error while reading file input:", err)
		return
	}

	words := reloaded.ProcessText(textContent)

	reloaded.WriteTextFile(outputFile, words)
}
