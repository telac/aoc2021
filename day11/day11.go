package day11

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	//"sort"
)

const nx = 10
const ny = 10

type Coords struct {
	x int
	y int
}

func readBoard(lines []string) []int {
	var board []int
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		for _, c := range line {
			ival, _ := strconv.Atoi(string(c))
			board = append(board, ival)
		}

	}
	return board
}

func isValid(x int, y int) bool {
	if x < 0 || x >= nx {
		return false
	}
	if y < 0 || y >= ny {
		return false
	}
	return true
}
func flashNeighbors(board []int, x int, y int) int {
	numFlashes := 0
	var neighbors = []Coords{
		Coords{x: x, y: y + 1},
		Coords{x: x, y: y - 1},
		Coords{x: x + 1, y: y},
		Coords{x: x - 1, y: y},
		Coords{x: x - 1, y: y - 1},
		Coords{x: x + 1, y: y + 1},
		Coords{x: x + 1, y: y - 1},
		Coords{x: x - 1, y: y + 1},
	}
	for _, coord := range neighbors {
		if isValid(coord.x, coord.y) {
			flashResult := incrementAndFlash(board, coord.x, coord.y)
			numFlashes += flashResult
			if flashResult == 1 {
				numFlashes += flashNeighbors(board, coord.x, coord.y)
			}
		}
	}
	return numFlashes
}

func incrementAndFlash(board []int, x int, y int) int {
	board[x+nx*y] += 1
	if board[x+nx*y] == 10 {
		return 1
	}
	return 0
}

func resetBoard(board []int) int {
	resetSquares := 0
	for i := 0; i < nx*ny; i++ {
		if board[i] > 9 {
			board[i] = 0
			resetSquares += 1
		}
	}
	return resetSquares
}

func evaluateTurn(board []int) int {
	numFlashes := 0
	for y := 0; y < ny; y++ {
		for x := 0; x < nx; x++ {
			numFlashes += incrementAndFlash(board, x, y)
			// flash neighbors
			if board[x+nx*y] == 10 {
				numFlashes += flashNeighbors(board, x, y)
			}
		}

	}
	return numFlashes
}

func visualizeBoard(b []int) {
	lineBuf := make([]int, nx)
	for i := 0; i < nx*ny; i++ {
		lineBuf[i%nx] = b[i]
		if i%nx == nx-1 {
			fmt.Println(lineBuf)
		}
	}
}

func runSim(board []int, steps int) int {
	numFlashes := 0
	for i := 0; i < steps; i++ {
		numFlashes += evaluateTurn(board)
		resetBoard(board)
	}
	return numFlashes
}

func syncSteps(board []int) int {
	resetCount := 0
	counter := 0
	for resetCount < 100 {
		counter += 1
		evaluateTurn(board)
		resetCount = resetBoard(board)
	}
	return counter
}

func Day11() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day11/input")
	lines := strings.Split(string(data), "\n")
	board := readBoard(lines)
	part1 := runSim(board, 100)
	part2 := syncSteps(board)
	return part1, part2 + 100
}
