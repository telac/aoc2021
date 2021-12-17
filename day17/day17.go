package day17

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Day17() (int, int) {
	pwd, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pwd + "/day17/example")
	lines := strings.Split(string(data), "\n")
	fmt.Println(lines)
	return 1, 0
}
