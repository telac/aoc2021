package day15

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// this language doesnt have a priorityqueue and im not writing one so im using this one: https://pkg.go.dev/container/heap
type Item struct {
	boardIndex int
	priority   int
	index      int
}

type Coords struct {
	x int
	y int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// return smallest
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type Board struct {
	board    []int
	distance []int
	visited  []bool
	items    []*Item
	nx       int
	ny       int
}

func readBoard(lines []string) Board {
	var board []int
	var distance []int
	var visited []bool
	var nx, ny int
	var items []*Item
	nx = len(lines[0])
	ny = len(lines) - 1
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		for _, c := range line {
			ival, _ := strconv.Atoi(string(c))
			board = append(board, ival)
			distance = append(distance, math.MaxInt64)
			visited = append(visited, false)
			items = append(items, nil)
		}

	}
	distance[0] = 0
	return Board{
		board:    board,
		distance: distance,
		visited:  visited,
		items:    items,
		nx:       nx,
		ny:       ny,
	}
}

func isValid(x int, y int, b Board) bool {
	if x < 0 || x >= b.nx {
		return false
	}
	if y < 0 || y >= b.ny {
		return false
	}
	return true
}

func getNeighborIndexes(b Board, index int) []int {
	x := index % b.nx
	y := index / b.ny
	var neighbors = []Coords{
		{x: x, y: y + 1},
		{x: x, y: y - 1},
		{x: x + 1, y: y},
		{x: x - 1, y: y},
	}
	var n []int
	for _, coord := range neighbors {
		if isValid(coord.x, coord.y, b) && !b.visited[coord.x+b.nx*coord.y] {
			n = append(n, coord.x+b.nx*coord.y)
		}
	}
	return n
}

func visualizeBoard(b Board) {
	fmt.Println()
	lineBuf := make([]int, b.nx)
	for i := 0; i < b.nx*b.ny; i++ {
		lineBuf[i%b.nx] = b.distance[i]
		if i%b.nx == b.nx-1 {
			fmt.Println(lineBuf)
		}
	}
}

func printPq(pq PriorityQueue) {
	for i := 0; i < pq.Len(); i++ {
		fmt.Println(pq[i])
	}
}

func popEverything(pq PriorityQueue) {
	fmt.Println("================")
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d  %.2d   %.2d \n", item.priority, item.boardIndex, item.index)
	}
}

func enlargeMap(b Board) Board {
	newBoard := make([]int, b.nx*b.ny*25)
	distance := make([]int, b.nx*b.ny*25)
	visited := make([]bool, b.nx*b.ny*25)
	items := make([]*Item, b.nx*b.ny*25)
	for yAxis := 0; yAxis < 5; yAxis++ {
		for xAxis := 0; xAxis < 5; xAxis++ {
			for i, v := range b.board {
				x := i % b.nx
				y := i / b.nx

				newVal := v + xAxis + yAxis
				if newVal > 9 {
					newVal = (newVal % 10) + 1
				}
				newX := x + (xAxis * b.nx)
				newY := y + (yAxis * b.ny)
				newBoard[newX+b.nx*5*newY] = newVal
				distance[newX+b.nx*5*newY] = math.MaxInt64
				visited[newX+b.nx*5*newY] = false
			}
		}
	}
	distance[0] = 0
	return Board{
		board:    newBoard,
		nx:       b.nx * 5,
		ny:       b.ny * 5,
		distance: distance,
		visited:  visited,
		items:    items,
	}

}

func dijsktra(b Board, start int, end int) Board {

	var pq PriorityQueue
	heap.Init(&pq)
	for k, v := range b.distance {
		item := &Item{
			priority:   v,
			boardIndex: k,
		}
		b.items[k] = item
		heap.Push(&pq, item)
	}

	counter := 0
	for pq.Len() > 0 {

		counter++
		current := heap.Pop(&pq).(*Item)

		if current.boardIndex == end {
			return b
		}

		b.visited[current.boardIndex] = true

		neighbors := getNeighborIndexes(b, current.boardIndex)

		for _, index := range neighbors {
			if !b.visited[index] {
				alt := current.priority + b.board[index]
				if alt < b.distance[index] {
					b.distance[index] = alt
					pq.update(b.items[index], alt)
				}
			}
		}
	}
	return b

}

func Day15() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day15/input")
	lines := strings.Split(string(data), "\n")
	board := readBoard(lines)
	b := dijsktra(board, 0, len(board.distance)-1)
	enlargedMap := enlargeMap(readBoard(lines))
	c := dijsktra(enlargedMap, 0, len(enlargedMap.distance)-1)
	return b.distance[len(b.distance)-1], c.distance[len(c.distance)-1]
}
