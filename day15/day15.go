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
	nx       int
	ny       int
}

func readBoard(lines []string) Board {
	var board []int
	var distance []int
	var visited []bool
	var nx, ny int
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
		}

	}
	distance[0] = 0
	return Board{
		board:    board,
		distance: distance,
		visited:  visited,
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
		if isValid(coord.x, coord.y, b) {
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
func Dijkstra(b Board, start int, end int) Board {
	//var path []int
	counter := 0
	var pq PriorityQueue
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		priority:   0,
		index:      0,
		boardIndex: 0,
	})
	b.distance[0] = 0
	for pq.Len() > 0 {
		counter += 1
		current := heap.Pop(&pq).(*Item)
		neighbors := getNeighborIndexes(b, current.boardIndex)
		b.visited[current.boardIndex] = true
		for _, index := range neighbors {
			if !b.visited[index] {
				alt := current.priority + b.board[index]
				heap.Push(&pq, &Item{
					priority:   alt,
					index:      index,
					boardIndex: index,
				})
				if alt < b.distance[index] {
					b.distance[index] = alt
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
	//fmt.Println(board)
	b := Dijkstra(board, 0, 99)
	return b.distance[len(b.distance)-1], 1
}
