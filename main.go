package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/internship_27_test/interfaces/commandline"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var result [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // 空行が来たら終了
		}
		words := strings.Split(line, " ")
		result = append(result, words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// panic recovery
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(os.Stderr, "Panic recovered: %v\n", r)
		}
	}()

	server := commandline.NewServer()
	err := server.Run(result)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
