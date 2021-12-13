package day13

import (
	"io/ioutil"
	"os"
	"fmt"
	"strconv"
	"strings"
	//"sort"
)

type Board struct {
	board []int
	nx int
	ny int
}

func readBoardAndInstructions(lines []string) (Board, []string) {
	var instructions []string
	var coordsList [][]int
	var nx int
	var ny int
	nx = 0
	ny = 0
	coordinatesEnd := false
	for _, line := range lines {
		if len(line) < 2 {
			coordinatesEnd = true
			continue
		}
		if coordinatesEnd {
			instructions = append(instructions, line)
		} else {
			coords := strings.Split(string(line), ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			if x + 1 > nx {
				nx = x + 1
			}
			if y + 1 > ny {
				ny = y + 1
			}
			coordsList = append(coordsList, []int{x, y})
		}

	}
	board := make([]int, nx*ny)
	for _, val := range coordsList {
		board[val[0] + nx * val[1]] = 1
	}

	b := Board{
		board : board,
		nx : nx,
		ny:ny,
	}
	return b, instructions
}

func fold(b Board, instruction string) (Board, int) {
	dots := 0
	ins := strings.Split(instruction, " ")
	fold := strings.Split(ins[2], "=")
	axis := fold[0]
	var nxNew, nyNew int
	magnitude, _ := strconv.Atoi(fold[1])
	switch axis {
	case "x":
		nxNew = b.nx / 2
		nyNew = b.ny
	case "y":
		nxNew = b.nx
		nyNew = b.ny / 2
	}
	//fmt.Println(nxNew, nyNew, nxNew*nyNew)
	boardNew := make([]int, nxNew*nyNew)
	for y := 0; y < nyNew; y++ {
		for x := 0; x < nxNew; x++ {
			switch axis {
			case "x":
				shift := (magnitude - x) + magnitude
				left := b.board[x + b.nx * y]
				right := b.board[shift + b.nx * y]
				if left == 1 || right == 1 {
					dots++
					boardNew[x + nxNew * y] = 1
				} else {
					boardNew[x + nxNew * y] = 0
				}
			case "y":
				shift := (magnitude - y) + magnitude
				up := b.board[x + b.nx * y]
				down := b.board[x + b.nx * shift]
				if up == 1 || down == 1 {
					dots++
					boardNew[x + nxNew * y] = 1
				} else {
					boardNew[x + nxNew * y] = 0
				}

			}
		}
	}
	return Board{
		board : boardNew,
		nx : nxNew,
		ny : nyNew,
	},
	dots

}

func visualizeBoard(b Board) {
	lineBuf := make([]string, b.nx)
	for i := 0; i < b.nx * b.ny; i++ {
		if b.board[i] == 1 {
			lineBuf[i % b.nx] = "#"
		} else {
			lineBuf[i % b.nx] = "."
		}
		
		if i % b.nx == b.nx - 1 {
			fmt.Println(lineBuf)
		}
	}
}


func FoldPaper(board Board, instructions []string) {
	b := board
	c := 0
	for _, instruction := range instructions {
		b, c = fold(b, instruction)
		visualizeBoard(b)
		fmt.Println(c)
	}
}
func Day13() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day13/input")
	lines := strings.Split(string(data), "\n")
	board, instructions := readBoardAndInstructions(lines)
	_, dots := fold(board, instructions[0])
	FoldPaper(board, instructions)
	return dots, 1
}
