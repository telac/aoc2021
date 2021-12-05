package day05

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
)

type Coordinates struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Board struct {
	nx int
	ny int
	coordinates []Coordinates
}

func trimAndCast(s string) int {
	val, _ := strconv.Atoi(strings.TrimSpace(s))
	return val
}

func getCoordinates(line string) Coordinates {
	pairs := strings.Split(line, "->")
	pair_1 := strings.Split(pairs[0], ",")
	pair_2 := strings.Split(pairs[1], ",")
	x1 := trimAndCast(pair_1[0])
	y1 := trimAndCast(pair_1[1])
	x2 := trimAndCast(pair_2[0])
	y2 := trimAndCast(pair_2[1])
	return Coordinates{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
	 

}

func buildBoard(lines []string) Board {
	var maxWidth = 0
	var maxDepth = 0
	var coords []Coordinates
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		coordinates := getCoordinates(line)

		if coordinates.x1 > maxWidth {
			maxWidth = coordinates.x1
		}
		if coordinates.x2 > maxWidth {
			maxWidth = coordinates.x2
		}
		if coordinates.y1 > maxDepth {
			maxDepth = coordinates.y1
		}
		if coordinates.y2 > maxDepth {
			maxDepth = coordinates.y2
		}
		coords = append(coords, coordinates)
	}
	return Board{
		nx: maxWidth + 1,
		ny: maxDepth + 1,
		coordinates: coords,
	}

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func buildLines(board Board, vertical bool) []int {
	lines := make([]int, board.nx * board.ny)
	for _, coords := range board.coordinates {
		horizontalStartpoint := min(coords.x1, coords.x2)
		horizontalEndpoint := max(coords.x1, coords.x2)
		verticalStartpoint := min(coords.y1, coords.y2)
		verticalEndpoint := max(coords.y1, coords.y2)
		horizontalDistance := horizontalEndpoint - horizontalStartpoint
		verticalDistance := verticalEndpoint - verticalStartpoint
		if horizontalDistance > 0 && verticalDistance > 0 && vertical  {
			startY := coords.y1
			endY := coords.y2
			if coords.x1 > coords.x2 {
				startY = coords.y2
				endY = coords.y1
			}
			upwards := startY < endY
			verticalPoint := startY
			for i := horizontalStartpoint; i <= horizontalEndpoint; i++ {
				lines[i + board.nx * verticalPoint] += 1
				if upwards {
					verticalPoint += 1
				} else {
					verticalPoint -= 1
				}
			}
		} else if horizontalDistance > 0 {

			for i := horizontalStartpoint; i <= horizontalEndpoint; i++ {
				lines[i + board.nx * coords.y1] += 1
			}
		} else {
			for i := verticalStartpoint; i <= verticalEndpoint; i++ {
				lines[coords.x1 + board.nx * i] += 1
			}
		}
	}
	return lines
}

func visualizeBoard(intersections []int, board Board) {
	lineBuf := make([]int, board.nx)
	for i := 0; i < board.nx * board.ny; i++ {
		lineBuf[i % board.nx] = intersections[i]
		if i % board.nx == 9 {
			fmt.Println(lineBuf)
		}
	}
}

func countPoints(intersections []int) int {
	points := 0
	for _, x := range intersections {
		if x > 1 {
			points++
		}
	}
	return points
}

func Day05() {
	pwd, _ := os.Getwd()
    data, _ := ioutil.ReadFile(pwd + "/day05/input")
	lines := strings.Split(string(data), "\n")
	board := buildBoard(lines)
	intersections1 := buildLines(board, false)
	points1 := countPoints(intersections1)
	//visualizeBoard(intersections1, board)
	fmt.Println(points1)
	intersections2 := buildLines(board, true)
	points2 := countPoints(intersections2)
	//visualizeBoard(intersections2, board)
	fmt.Println(points2)


}
