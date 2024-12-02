package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readReports()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	safeReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
			fmt.Println(report)
		}
	}
	fmt.Println(safeReports)
	safeReports = 0
	for _, report := range reports {
		if isReportSafeWithDampener(report) {
			safeReports++
			fmt.Println(report)
		}
	}
	fmt.Println(safeReports)
}

func isReportSafe(report []int) bool {
	for i := range report {
		if i == 0 {
			continue
		}
		safe := isLevelSave(report, i, report[1] > report[0])
		if !safe {
			return false
		}
	}
	return true
}

func isReportSafeWithDampener(report []int) bool {
	for i := range report {
		clonedReport := make([]int, len(report))
		copy(clonedReport, report)
		if isReportSafe(remove(clonedReport, i)) {
			return true
		}
	}

	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isLevelSave(report []int, i int, isIncreasing bool) bool {
	lastLevel := report[i-1]
	actuallyLevel := report[i]
	if lastLevel < actuallyLevel {
		if !isIncreasing {
			return false
		}
		diff := actuallyLevel - lastLevel
		if diff < 1 || diff > 3 {
			return false
		}
	} else if lastLevel > actuallyLevel {
		if isIncreasing {
			return false
		}
		diff := lastLevel - actuallyLevel
		if diff < 1 || diff > 3 {
			return false
		}
	} else {
		return false
	}
	return true
}

func readReports() ([][]int, error) {
	var reports [][]int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input (press Ctrl+D or Ctrl+Z to end):")
	reportNumber := 0
	for scanner.Scan() {
		reportLine := scanner.Text()
		levels := strings.Fields(reportLine)
		reports = append(reports, make([]int, len(levels)))
		for levelIndex, level := range levels {
			levelNumber, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			reports[reportNumber][levelIndex] = levelNumber
		}
		reportNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return reports, nil
}
