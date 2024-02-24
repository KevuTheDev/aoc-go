package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	content, err := os.Open("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("End of File")

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		token := scanner.Text()
	}
}

func part2() {
	content, err := os.Open("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("End of File")

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		token := scanner.Text()
	}
}

func main() {
	switcher := true

	if switcher {
		part1()
	} else {
		part2()
	}
}
