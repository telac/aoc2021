package day16

import (
	"io/ioutil"
	"math"
	"os"
	"strings"

	//"strings"
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
)

type Packet struct {
	uuid    string
	version int64
	id      int64
	tp      string
	parent  *Packet
	depth   int64
	value   int64
	length  int64
}

func readPackageMap(elfCode string, offset int64, depth int64, parent *Packet) ([]*Packet, map[Packet][]*Packet, int64) {
	var packets []*Packet

	edges := make(map[Packet][]*Packet)
	lengths := map[int64]int{
		0: 15,
		1: 11,
	}
	id, _ := strconv.ParseInt(elfCode[offset+3:offset+6], 2, 64)
	switch id {
	case 4:
		packet := readLiteral(elfCode, offset, depth, parent)
		offset += packet.length
		packets = append(packets, &packet)
		edges[packet] = append(edges[packet], nil)
		return packets, edges, offset
	default:
		version, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
		offset += 3
		id, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
		offset += 3
		lengthType, _ := strconv.ParseInt(elfCode[offset:offset+1], 2, 64)
		offset += 1
		subOperatorLength := int64(lengths[lengthType])
		msgLength, _ := strconv.ParseInt(elfCode[offset:offset+subOperatorLength], 2, 64)
		offset += subOperatorLength
		oldOffset := offset
		switch lengthType {
		case 0:
			//you basically need luck to get the solution right, but i got it after a few attempts. a GUID would be cool i guess.
			u, _ := uuid.NewV4()
			cPacket := &Packet{
				version: version,
				id:      id,
				value:   -1,
				tp:      "operator",
				uuid:    u.String(),
				depth:   depth,
				parent:  parent,
				length:  7 + subOperatorLength,
			}
			packets = append(packets, cPacket)
			for offset < (msgLength + oldOffset) {
				newPackets, _, newOffset := readPackageMap(elfCode, offset, depth+1, cPacket)
				offset = newOffset
				packets = append(packets, newPackets...)

			}
		case 1:
			u, _ := uuid.NewV4()
			cPacket := &Packet{
				version: version,
				id:      id,
				value:   -1,
				tp:      "operator",
				uuid:    u.String(),
				depth:   depth,
				parent:  parent,
				length:  7 + subOperatorLength,
			}
			packets = append(packets, cPacket)
			for i := 0; i < int(msgLength); i++ {
				newPackets, _, newOffset := readPackageMap(elfCode, offset, depth+1, cPacket)
				offset = newOffset
				packets = append(packets, newPackets...)

			}

		}
	}
	return packets, edges, offset
}

func makeEdges(packets []*Packet) map[Packet][]*Packet {
	edges := make(map[Packet][]*Packet)
	for _, v := range packets {
		if v.parent != nil {
			edges[*v.parent] = append(edges[*v.parent], v)
		}
	}
	return edges
}

func handleOperation(packetAddr *Packet, edges map[Packet][]*Packet) int64 {
	var value int64
	packet := *packetAddr
	//fmt.Printf("%p \n", packetAddr)
	switch packet.id {
	case 0:
		var sum int64
		sum = 0
		var values []int64
		for _, v := range edges[packet] {
			temp := handleOperation(v, edges)
			values = append(values, temp)
			sum += temp
		}
		value = sum
	case 1:
		var product int64
		product = 1
		for _, v := range edges[packet] {
			tmp := handleOperation(v, edges)
			product *= tmp
		}
		value = product
	case 2:
		var min int64
		min = math.MaxInt64
		for _, v := range edges[packet] {
			tMin := handleOperation(v, edges)
			if tMin < min {
				min = tMin
			}
		}
		value = min
	case 3:
		var max int64
		max = 0
		for _, v := range edges[packet] {
			tmax := handleOperation(v, edges)
			if tmax > max {
				max = tmax
			}
		}
		value = max
	case 4:
		value = packet.value
	case 5:

		val1 := handleOperation(edges[packet][0], edges)
		val2 := handleOperation(edges[packet][1], edges)
		if val1 > val2 {
			value = 1
		} else {
			value = 0
		}

	case 6:
		val1 := handleOperation(edges[packet][0], edges)
		val2 := handleOperation(edges[packet][1], edges)
		if val1 < val2 {
			value = 1
		} else {
			value = 0
		}

	case 7:
		val1 := handleOperation(edges[packet][0], edges)
		val2 := handleOperation(edges[packet][1], edges)
		if val1 == val2 {
			value = 1
		} else {
			value = 0
		}
	}
	return value
}

func readLiteral(elfCode string, offset int64, depth int64, parent *Packet) Packet {
	origOffset := offset
	version, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
	offset += 3
	id, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
	offset += 3
	end := false
	var bitValue string
	for !end {
		if string(elfCode[offset]) == "0" {
			end = true
		}
		offset++
		bitValue += elfCode[offset : offset+4]
		offset += 4
	}
	iBitvalue, _ := strconv.ParseInt(bitValue, 2, 64)
	u, _ := uuid.NewV4()
	return Packet{
		version: version,
		id:      id,
		value:   iBitvalue,
		tp:      "literal",
		uuid:    u.String(),
		depth:   depth,
		parent:  parent,
		length:  offset - origOffset,
	}
}

func countVersions(packets []*Packet) int {
	var versionCount int64
	for _, v := range packets {
		versionCount += v.version
	}
	return int(versionCount)
}

func toBinary(line string) string {
	characters := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	var newString string
	for _, r := range line {
		newString += characters[string(r)]
	}
	return newString
}

func Day16() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day16/input")
	lines := strings.Split(string(data), "\n")
	binaryValue := toBinary(lines[0])
	newpackets, _, _ := readPackageMap(binaryValue, 0, 0, nil)
	newEdges := makeEdges(newpackets)
	part2 := handleOperation(newpackets[0], newEdges)
	return countVersions(newpackets), int(part2)
}
