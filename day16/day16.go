package day16

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	//"strings"
	"strconv"
)

type Packet struct {
	version int64
	id      int64
	tp      string
	value   int64
	length  int64
}

func readPackage(elfCode string, offset int64) ([]Packet, int64) {
	lengths := map[int64]int{
		0: 15,
		1: 11,
	}
	id, _ := strconv.ParseInt(elfCode[offset+3:offset+6], 2, 64)
	//fmt.Println("found version|id", elfCode[offset:])
	//fmt.Println("offset is now:", offset)
	var packets []Packet
	switch id {
	case 4:
		packet := readLiteral(elfCode, offset)
		offset += packet.length
		packets = append(packets, packet)
		return packets, offset
	default:
		version, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
		offset += 3
		id, _ := strconv.ParseInt(elfCode[offset:offset+3], 2, 64)
		offset += 3
		lengthType, _ := strconv.ParseInt(elfCode[offset:offset+1], 2, 64)
		offset += 1
		subOperatorLength := int64(lengths[lengthType])
		msgLength, _ := strconv.ParseInt(elfCode[offset:offset+subOperatorLength], 2, 64)
		//fmt.Println(elfCode[offset : offset+subOperatorLength])
		offset += subOperatorLength
		oldOffset := offset
		switch lengthType {
		case 0:
			packets = append(packets, Packet{
				version: version,
				id:      id,
				value:   -1,
				tp:      "operator",
				length:  7 + subOperatorLength,
			})
			fmt.Println(offset, msgLength, subOperatorLength)
			for offset < (msgLength + oldOffset) {
				newPackets, newOffset := readPackage(elfCode, offset)
				offset = newOffset
				packets = append(packets, newPackets...)
			}
		case 1:
			packets = append(packets, Packet{
				version: version,
				id:      id,
				value:   -1,
				tp:      "operator",
				length:  7 + subOperatorLength,
			})
			fmt.Println("reading total of ", msgLength, "packages")
			for i := 0; i < int(msgLength); i++ {
				newPackets, newOffset := readPackage(elfCode, offset)
				offset = newOffset
				packets = append(packets, newPackets...)
			}
		}
	}
	return packets, offset
}

func readLiteral(elfCode string, offset int64) Packet {
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
	return Packet{
		version: version,
		id:      id,
		value:   iBitvalue,
		tp:      "literal",
		length:  offset - origOffset,
	}
}

func countVersions(packets []Packet) int {
	var versionCount int64
	for _, v := range packets {
		versionCount += v.version
	}
	return int(versionCount)
}

func getPackages(elfCode string) []Packet {
	packets, _ := readPackage(elfCode, 0)
	fmt.Println(packets)
	return packets
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
	data, _ := ioutil.ReadFile(pwd + "/day16/example")
	lines := strings.Split(string(data), "\n")
	binaryValue := toBinary(lines[0])
	fmt.Println(binaryValue)
	packets := getPackages(binaryValue)
	return countVersions(packets), 0
}
