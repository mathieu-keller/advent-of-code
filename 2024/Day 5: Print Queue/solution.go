package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pageRules, pages, err := parseInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	goodPages, badPages := classifyPages(pages, pageRules)

	// Calculate middle element sum for good pages
	middleSum := calculateMiddleSum(goodPages)
	fmt.Println("Good Pages Middle Sum:", middleSum)

	// Correct bad pages and calculate their middle element sum
	correctedPages := correctBadPages(badPages, pageRules)
	middleSum = calculateMiddleSum(correctedPages)
	fmt.Println("Corrected Pages Middle Sum:", middleSum)
}

func classifyPages(pages [][]int, rules map[int][]int) ([][]int, [][]int) {
	var goodPages, badPages [][]int
	for _, page := range pages {
		if isPageValid(page, rules) {
			goodPages = append(goodPages, page)
		} else {
			badPages = append(badPages, page)
		}
	}
	return goodPages, badPages
}

func isPageValid(page []int, rules map[int][]int) bool {
	for i, value := range page {
		if rule, exists := rules[value]; exists && containsAny(page[:i], rule) {
			return false
		}
	}
	return true
}

func correctBadPages(badPages [][]int, rules map[int][]int) [][]int {
	var correctedPages [][]int
	for _, page := range badPages {
		correctedPages = append(correctedPages, correctPage(page, rules))
	}
	return correctedPages
}

func correctPage(page []int, rules map[int][]int) []int {
	for i, value := range page {
		if rule, exists := rules[value]; exists && containsAny(page[:i], rule) {
			targetIndex := findIndex(page, value)
			page = moveElement(page, targetIndex, i-1)
			return correctPage(page, rules)
		}
	}
	return page
}

func calculateMiddleSum(pages [][]int) int {
	middleSum := 0
	for _, page := range pages {
		if len(page) > 0 {
			middleSum += page[len(page)/2]
		}
	}
	return middleSum
}

func parseInput() (map[int][]int, [][]int, error) {
	fmt.Println("Enter input (press Ctrl+D to finish):")
	scanner := bufio.NewScanner(os.Stdin)

	rules := make(map[int][]int)
	var pages [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			key, value, err := parseRule(line)
			if err != nil {
				return nil, nil, err
			}
			rules[key] = append(rules[key], value)
		} else if strings.Contains(line, ",") {
			page, err := parsePage(line)
			if err != nil {
				return nil, nil, err
			}
			pages = append(pages, page)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, pages, nil
}

func parseRule(line string) (int, int, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid rule format: %s", line)
	}

	key, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return key, value, nil
}

func parsePage(line string) ([]int, error) {
	parts := strings.Split(line, ",")
	page := make([]int, len(parts))
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		page[i] = value
	}
	return page, nil
}

func moveElement(arr []int, fromIndex, toIndex int) []int {
	if fromIndex == toIndex {
		return arr
	}

	element := arr[fromIndex]
	arr = append(arr[:fromIndex], arr[fromIndex+1:]...)
	return append(arr[:toIndex], append([]int{element}, arr[toIndex:]...)...)
}

func containsAny(slice, elements []int) bool {
	elementSet := make(map[int]struct{}, len(elements))
	for _, e := range elements {
		elementSet[e] = struct{}{}
	}
	for _, s := range slice {
		if _, found := elementSet[s]; found {
			return true
		}
	}
	return false
}

func findIndex(slice []int, target int) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1
}
