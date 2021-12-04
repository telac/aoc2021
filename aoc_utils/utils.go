package utils

func ReadLines() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 1 {
            if line[0] == '\n' {
                break
            }
        }
        iNum, err := strconv.Atoi(line)
        if err != nil {
           fmt.Println(err)
        }
        nums = append(nums, iNum)
    }
    return nums
}

func ReadStrLines() []str {
    var lines []str
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 1 {
            if line[0] == '\n' {
                break
            }
        }
        nums = append(nums, line)
    }
    return lines
}