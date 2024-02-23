package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	content, err := os.Open("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("End of File")

	scanner := bufio.NewScanner(content)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	counter := 0

	currentCalories := 0
	maxCalories := currentCalories

	caloriesArray := []int{}

	for scanner.Scan() {
		token := scanner.Text()
		if token == "" {
			caloriesArray = append(caloriesArray, currentCalories)
			if currentCalories > maxCalories {
				maxCalories = currentCalories

			}
			currentCalories = 0
		} else {
			i, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Error converting token to integer")
				panic(err)
			}
			currentCalories += i
		}
		fmt.Println(token)
		counter += 1
	}

	fmt.Println()
	fmt.Printf("Total count: %d\n", counter)
	fmt.Printf("Max calories: %d\n", maxCalories)

	sort.Ints(caloriesArray)
	caloriesArray = caloriesArray[len(caloriesArray)-3:]

	fmt.Printf("Last three Calories: %d\n", caloriesArray)
	sum := 0
	for i := 0; i < len(caloriesArray); i++ {
		sum += caloriesArray[i]
	}

	fmt.Printf("Sum of Last three Calories: %d\n", sum)
}
