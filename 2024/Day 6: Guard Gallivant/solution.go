package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

type guardState struct {
	pos position
	dir int
}

var directions = []rune{'^', '>', 'v', '<'}

func main() {
	original, err := parseInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	grid := make([][]byte, len(original))
	for i := range original {
		grid[i] = []byte(original[i])
	}

	startPos, startDir := findInitialPositionAndDirection(grid)
	visited := simulateGuardPath(grid, startPos, startDir)
	fmt.Println("Visited Fields:", len(visited))

	obstructionPositions := findObstructionPositions(grid, startPos, startDir)
	fmt.Println("Obstruction Positions:", len(obstructionPositions))
}

func parseInput() ([][]byte, error) {
	fmt.Println("Enter grid input (press Ctrl+D to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]byte

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(grid) == 0 {
		return nil, errors.New("no input provided")
	}

	return grid, nil
}

func findInitialPositionAndDirection(grid [][]byte) (position, int) {
	for y, row := range grid {
		for x, char := range row {
			if idx := strings.IndexRune("^>v<", rune(char)); idx != -1 {
				return position{x: x, y: y}, idx
			}
		}
	}
	panic("No guard found in the input grid")
}

func simulateGuardPath(grid [][]byte, startPos position, startDir int) map[position]bool {
	visited := make(map[position]bool)
	currentPos, currentDir := startPos, startDir

	for {
		visited[currentPos] = true
		dx, dy := getDirectionDeltas(currentDir)
		next := position{x: currentPos.x + dx, y: currentPos.y + dy}

		if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
			break
		}

		if grid[next.y][next.x] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			currentPos = next
		}
	}

	return visited
}

func findObstructionPositions(grid [][]byte, startPos position, startDir int) []position {
	var obstructionPositions []position

	for y, row := range grid {
		for x, cell := range row {
			pos := position{x: x, y: y}

			if pos == startPos || cell == '#' {
				continue
			}

			original := grid[y][x]
			grid[y][x] = '#'

			if isGuardTrapped(grid, startPos, startDir) {
				obstructionPositions = append(obstructionPositions, pos)
			}

			grid[y][x] = original
		}
	}

	return obstructionPositions
}

func isGuardTrapped(grid [][]byte, startPos position, startDir int) bool {
	currentPos, currentDir := startPos, startDir
	seenStates := make(map[guardState]bool)

	for {
		state := guardState{pos: currentPos, dir: currentDir}
		if seenStates[state] {
			return true
		}
		seenStates[state] = true

		dx, dy := getDirectionDeltas(currentDir)
		next := position{x: currentPos.x + dx, y: currentPos.y + dy}

		if next.y < 0 || next.y >= len(grid) || next.x < 0 || next.x >= len(grid[0]) {
			break
		}

		if grid[next.y][next.x] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			currentPos = next
		}
	}

	return false
}

func getDirectionDeltas(dir int) (int, int) {
	switch directions[dir] {
	case '^':
		return 0, -1
	case '>':
		return 1, 0
	case 'v':
		return 0, 1
	case '<':
		return -1, 0
	}
	panic("Invalid direction")
}
