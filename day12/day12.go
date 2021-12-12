package day12

import (
	"io/ioutil"
	"os"

	//"strconv"
	"strings"
	//"sort"
)

type Node struct {
	value string
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	start *Node
	end   *Node
}

func buildGraph(lines []string) Graph {
	seenNodes := make(map[string]Node)
	var nodes []*Node
	nodeMap := make(map[Node][]*Node)
	g := Graph{
		nodes: nodes,
		edges: nodeMap,
	}
	for _, v := range lines {
		if len(v) < 2 {
			continue
		}
		graphPoints := strings.Split(v, "-")
		startString := graphPoints[0]
		endString := graphPoints[1]
		var start, end Node
		_, oks := seenNodes[startString]
		if !oks {
			seenNodes[startString] = Node{startString}
			start := seenNodes[startString]
			if start.value == "start" {
				g.start = &start
			}
			if start.value == "end" {
				g.end = &end
			}
			g.nodes = append(g.nodes, &start)
		}
		_, oke := seenNodes[endString]
		if !oke {
			seenNodes[endString] = Node{endString}
			end = seenNodes[endString]
			g.nodes = append(g.nodes, &end)
			if end.value == "end" {
				g.end = &end
			}
			if end.value == "start" {
				g.start = &end
			}
		}
		start = seenNodes[startString]
		end = seenNodes[endString]

		g.edges[start] = append(g.edges[start], &end)
		g.edges[end] = append(g.edges[end], &start)

	}
	return g
}
func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].value + " <-> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].value + " "
		}
		s += "\n"
	}
	//fmt.Println(s)
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

var count = 0

func findPath(start *Node, end *Node, g Graph, visitedNodes map[Node]bool, path []string) int {
	count := 0
	visitedNodes[*start] = true
	path = append(path, start.value)
	if start.value == end.value {
		count++
	} else {
		for _, next := range g.edges[*start] {
			if _, ok := visitedNodes[*next]; !ok {
				visitedNodes[*next] = false
			}
			if isUpper(next.value) || !visitedNodes[*next] {
				copyVisits := make(map[Node]bool)
				for k, v := range visitedNodes {
					copyVisits[k] = v
				}
				//fmt.Print(next.value, end.value, copyVisits, path)
				count += findPath(next, end, g, copyVisits, path)
			}

		}
	}
	return count
}

func incrementMapValueByOne(m map[Node]int, k Node) map[Node]int {
	if _, ok := m[k]; !ok {
		m[k] = 1
	} else {
		m[k] += 1
	}
	return m
}

func findPathTwice(start *Node, end *Node, g Graph, visitedNodes map[Node]int, visitedTwice bool, path []string) int {
	count := 0
	incrementMapValueByOne(visitedNodes, *start)
	if visitedNodes[*start] == 2 && !isUpper(start.value) {
		visitedTwice = true
	}
	path = append(path, start.value)
	if start.value == end.value {
		count++
	} else {
		for _, next := range g.edges[*start] {
			if _, ok := visitedNodes[*next]; !ok {
				visitedNodes[*next] = 0
			}
			if next.value != "start" && (isUpper(next.value) || visitedNodes[*next] < 1 || !visitedTwice) {
				copyVisits := make(map[Node]int)
				for k, v := range visitedNodes {
					copyVisits[k] = v
				}
				count += findPathTwice(next, end, g, copyVisits, visitedTwice, path)
			}

		}
	}
	return count
}

func Day12() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day12/input")
	lines := strings.Split(string(data), "\n")
	g := buildGraph(lines)
	visited := make(map[Node]bool)
	visitCount := make(map[Node]int)
	var p []string
	p1 := findPath(g.start, g.end, g, visited, p)
	p2 := findPathTwice(g.start, g.end, g, visitCount, false, p)
	g.String()
	return p1, p2
}
