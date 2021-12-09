package day09

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sort"
)
type Board struct {
	ny int
	nx int
	data []int
}

type LowPoint struct {
	value int
	x int
	y int
}

func getData(lines []string) Board {
	nx := len(strings.Split(lines[0], ""))
	ny := len(lines) - 1
	data := make([]int, nx*ny)
	for i, v := range lines {
		for j, v := range v {
			convValue, _ := strconv.Atoi(string(v))
			data[j + nx * i] = convValue
		}
	}
	_board := Board{
		ny : ny,
		nx : nx,
		data : data,
	}
	return _board
}

func findLowPoints(b Board) []LowPoint {
	var lowPoints []LowPoint
	nx := b.nx
	ny := b.ny
	data := b.data
	for i := 0; i < b.ny; i++ {
		for j := 0; j < b.nx; j++ {
			var isLowPoint bool
			isLowPoint = true
			if i > 0 && data[j + nx * i] >= data[j + nx * (i - 1)] {
				isLowPoint = false
			}
			if i < ny - 1 && data[j + nx * i] >= data[j + nx * (i + 1)] {
				isLowPoint = false
			}
			if j > 0 && data[j + nx * i] >= data[(j - 1) + nx * i] {
				isLowPoint = false
			}
			if j < nx - 1 && data[j + nx * i] >= data[(j + 1) + nx * i] {
				isLowPoint = false
			}
			if isLowPoint {		
				lowPoint := LowPoint {
					x : j,
					y : i,
					value : data[j + nx * i],
				}
				lowPoints = append(lowPoints, lowPoint)
			}
		}
	}
	return lowPoints
}

func getNeighbors(coord int, b Board, visited map[int]int) []int {
	x := coord % b.nx
	y := coord / b.nx
	data := b.data
	var neighbors []int
	if x > 0 && data[x - 1 + b.nx * y] < 9 {
		index := x - 1 + b.nx * y
		_, ok := visited[index]
		if !ok {
		neighbors = append(neighbors, index)
		}
	}
	if x < b.nx - 1 && data[x + 1 + b.nx * y] < 9 {
		neighbors = append(neighbors, x + 1 + b.nx * y)
	}
	if y < b.ny - 1 && data[x+ b.nx * (y + 1)] < 9 {
		neighbors = append(neighbors, x+ b.nx * (y + 1))
	}
	if y > 0 && data[x+ b.nx * (y - 1)] < 9 {
		neighbors = append(neighbors, x+ b.nx * (y - 1))
	}
	return neighbors
}


func findBasin(lp LowPoint, b Board) map[int]int {
	visited := make(map[int]int)
	coord1D := lp.x + b.nx * lp.y
	var currentNode int
	var neighbors []int
	neighbors = append(neighbors, coord1D)
	visited[coord1D] = b.data[coord1D]
	for len(neighbors) > 0 {
		currentNode, neighbors = neighbors[len(neighbors)-1], neighbors[:len(neighbors)-1]
		newCandidates := getNeighbors(currentNode, b, visited)
		for _, v := range newCandidates {
			_, ok := visited[v]
			if !ok{
				visited[v] = b.data[v]
				neighbors = append(neighbors, v)
			}
		}
	}
	return visited

}

func findBasins(lp []LowPoint, b Board) int {
	var basinsizes []int
	for _, v := range lp {
		basin := findBasin(v, b)
		basinsizes = append(basinsizes, len(basin))
	}
	sort.Slice(basinsizes, func(a, b int) bool {
		return basinsizes[b] < basinsizes[a]
	 })
	return basinsizes[0] * basinsizes[1] * basinsizes[2]
}

func sumLowPoints(a []LowPoint) int {
	sum := 0
	for _, v := range a {
		sum += v.value
		sum += 1
	}
	return sum
}

func visualizeBoard(b Board) {
	lineBuf := make([]int, b.nx)
	for i := 0; i < b.nx*b.ny; i++ {
		lineBuf[i%b.nx] = b.data[i]
		if i%b.nx == b.nx-1 {
			fmt.Println(lineBuf)
		}
	}
}

func Day09() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day09/input")
	lines := strings.Split(string(data), "\n")
	b := getData(lines)
	
	lowPoints := findLowPoints(b)
	basinCount := findBasins(lowPoints, b)
	//visualizeBoard(b)
	return sumLowPoints(lowPoints), basinCount
}
