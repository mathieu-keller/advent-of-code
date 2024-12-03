package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	memory, _ := readMemory()
	calculateOnlyActives(memory)
	calculateAll(memory)
}

func calculateOnlyActives(memory string) {
	activeCommands := getActiveCommands(memory)
	result := 0
	for _, activeCommand := range activeCommands {
		mulInput := getMulCommands(activeCommand)
		result += calculate(mulInput)
	}
	fmt.Println("active calculation: " + strconv.Itoa(result))
}

func calculateAll(memory string) {
	mulInput := getMulCommands(memory)
	fmt.Println("all calculation: " + strconv.Itoa(calculate(mulInput)))
}

func calculate(mulInput []string) int {
	result := 0
	for _, input := range mulInput {
		mulResult, err := mul(input)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		result += mulResult
	}
	return result
}

func getMulCommands(memory string) []string {
	startingInput := strings.Split(memory, "mul(")
	mulInput := make([]string, 0)
	for _, start := range startingInput {
		if strings.Contains(start, ")") {
			endInput := strings.Split(start, ")")
			match, _ := regexp.MatchString("^\\d*,\\d*$", endInput[0])
			if match {
				mulInput = append(mulInput, endInput[0])
			}
		}
	}
	return mulInput
}

func getActiveCommands(memory string) []string {
	startDoes := strings.Split(memory, "do()")
	does := make([]string, len(startDoes))
	for i, startDo := range startDoes {
		does[i] = strings.Split(startDo, "don't")[0]
	}
	return does
}

func mul(input string) (int, error) {
	inputSplit := strings.Split(input, ",")
	a, err := strconv.Atoi(inputSplit[0])
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(inputSplit[1])
	if err != nil {
		return 0, err
	}
	return a * b, nil
}

func readMemory() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input (press Ctrl+D or Ctrl+Z to end):")
	memory := ""
	for scanner.Scan() {
		reportLine := scanner.Text()
		memory += reportLine
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return memory, nil
}
