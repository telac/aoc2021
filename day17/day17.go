package day17

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Target struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
}

type Coords struct {
	x int
	y int
}

func getTarget(line string) Target {
	parts := strings.Split(line, " ")
	fmt.Println(parts)
	x := strings.Trim(strings.Split(parts[2], "=")[1], ",")
	xStart, _ := strconv.Atoi(strings.Split(x, "..")[0])
	xEnd, _ := strconv.Atoi(strings.Split(x, "..")[1])
	y := strings.Split(parts[3], "=")[1]
	yStart, _ := strconv.Atoi(strings.Split(y, "..")[0])
	yEnd, _ := strconv.Atoi(strings.Split(y, "..")[1])
	return Target{
		xStart: xStart,
		xEnd:   xEnd,
		yStart: yStart,
		yEnd:   yEnd,
	}
}

func simTrajectory(x int, y int, target Target) (int, int) {
	xPos := 0
	yPos := 0
	hitCount := 0
	maxY := 0
	for yPos >= target.yStart {
		xPos += x
		yPos += y
		if x > 0 {
			x--
		}
		if yPos > maxY {
			maxY = yPos
		}
		y--
		if xPos >= target.xStart && xPos <= target.xEnd && yPos >= target.yStart && yPos <= target.yEnd {
			hitCount++

		}
	}
	return hitCount, maxY
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
func minMax(target Target) Coords {
	y0 := abs(target.yStart)
	y1 := abs(target.yEnd)
	fmt.Println(y0, y1)
	if y0 > y1 {
		return Coords{
			x: y1,
			y: y0,
		}
	}
	return Coords{
		x: y0,
		y: y1,
	}
}

func getHighestPosition(target Target) (Coords, int, int) {
	maxY := 0
	numCoordinates := 0
	var coords Coords
	//yRange := minMax(target)
	for x := 0; x <= target.xEnd; x++ {
		for y := -129; y <= 129; y++ {
			//fmt.Println(x, y)
			hits, topY := simTrajectory(x, y, target)
			if hits > 0 {
				numCoordinates++
				if topY > maxY {
					maxY = topY
					coords = Coords{
						x: x,
						y: y,
					}
				}
			}
		}
	}
	return coords, maxY, numCoordinates
}

func Day17() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day17/input")
	lines := strings.Split(string(data), "\n")
	target := getTarget(lines[0])
	_, part1, part2 := getHighestPosition(target)
	return part1, part2
}
