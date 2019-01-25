package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	testScanner()
}

func testScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		upper := strings.ToUpper(scanner.Text())
		fmt.Println(upper)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
