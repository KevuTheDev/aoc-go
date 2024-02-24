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

	total := 0
	for scanner.Scan() {
		token := scanner.Text()
		m := make(map[string]int)

		half1 := token[:len(token)/2]
		half2 := token[len(token)/2:]

		for i := 0; i < len(half1); i++ {
			tk := string(half1[i])
			_, ok := m[tk]

			if !ok {
				m[tk] = 0
			}
		}

		for i := 0; i < len(half2); i++ {
			tk := string(half2[i])
			_, ok := m[tk]

			if ok {
				m[tk] += 1
			}
		}

		for key, val := range m {
			if val == 0 {
				delete(m, key)
			} else {
				value := int([]rune(key)[0])
				if value >= 65 && value <= 90 {
					fmt.Println(value - 38)
					total += value - 38
				} else {
					fmt.Println(value - 96)
					total += value - 96
				}
			}
		}

		//fmt.Println(m)

		//fmt.Printf("1: %s <+++> %s :2\n", half1, half2)
	}
	fmt.Printf("Total: %d", total)
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
		fmt.Println(token)
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
