package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2, err := readList()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(list1) != len(list2) {
		fmt.Println("lists have not the same length!")
		os.Exit(1)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	fmt.Println("The different is: " + strconv.Itoa(calculateDifference(list1, list2)))
	fmt.Println("The score is: " + strconv.Itoa(calculateScore(list1, list2)))
}

func calculateScore(list1 []int, list2 []int) int {
	result := 0
	resultMap := make(map[int]int)
	for _, number1 := range list1 {
		_, ok := resultMap[number1]
		if !ok {
			appears := 0
			for _, number2 := range list2 {
				if number1 == number2 {
					appears++
				}
			}
			resultMap[number1] = number1 * appears
		}
		result += resultMap[number1]

	}
	return result
}

func calculateDifference(list1 []int, list2 []int) int {
	result := 0
	for i, number1 := range list1 {
		number2 := list2[i]
		if number1 >= number2 {
			result += number1 - number2
		} else {
			result += number2 - number1
		}
	}
	return result
}

func readList() ([]int, []int, error) {
	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input (press Ctrl+D or Ctrl+Z to end):")

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			list1Num, err1 := strconv.Atoi(parts[0])
			list2Num, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, list1Num)
				list2 = append(list2, list2Num)
			} else {
				fmt.Println("Invalid input. Please enter two integers separated by spaces.")
			}
		} else {
			fmt.Println("Invalid input format. Each line should have exactly two numbers.")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return list1, list2, nil
}
