package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for {
		if !scanner.Scan() {
			continue
		}
		fmt.Println(scanner.Text())
	}
}
