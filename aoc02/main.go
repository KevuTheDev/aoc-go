package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	content, err := os.Open("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("End of File")

	scanner := bufio.NewScanner(content)

	playerPoints := 0

	for scanner.Scan() {
		token := scanner.Text()
		if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 1
			playerPoints += 3
		} else if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 2
			playerPoints += 6
		} else if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 3
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 1
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 2
			playerPoints += 3
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 3
			playerPoints += 6
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 1
			playerPoints += 6
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 2
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 3
			playerPoints += 3
		}
	}

	fmt.Printf("Player has %d points!\n", playerPoints)
}

func part2() {
	content, err := os.Open("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("End of File")

	scanner := bufio.NewScanner(content)

	playerPoints := 0

	for scanner.Scan() {
		token := scanner.Text()
		if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 3
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 1
			playerPoints += 3
		} else if strings.Compare(string(token[0]), "A") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 2
			playerPoints += 6
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 1
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 2
			playerPoints += 3
		} else if strings.Compare(string(token[0]), "B") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 3
			playerPoints += 6
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "X") == 0 {
			playerPoints += 2
			playerPoints += 0
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "Y") == 0 {
			playerPoints += 3
			playerPoints += 3
		} else if strings.Compare(string(token[0]), "C") == 0 && strings.Compare(string(token[2]), "Z") == 0 {
			playerPoints += 1
			playerPoints += 6
		}
	}

	fmt.Printf("Player has %d points!\n", playerPoints)
}

func main() {
	part1()
	part2()
}
