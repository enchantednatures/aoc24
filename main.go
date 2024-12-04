package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a new instance of the server
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Data: %s\n", data)

	strData := string(data)

	l, r := make([]int, 0), make([]int, 0)

	lines := strings.Split(strings.TrimSpace(strData), "\n")
	for lineNum, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " ")
		for idx, v := range parts {
			fmt.Printf("Part: %s, %d\n", v, idx)
		}

		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error parsing left value on line %d: %v", lineNum, err)
			return
		}
		rightVal, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {

			log.Fatalf("Error parsing right value on line %d: %v", lineNum, err)
			return
		}
		l = append(l, leftVal)
		r = append(r, rightVal)

	}

	val := process2(l, r)
	fmt.Printf("Value: %d\n", val)
}

func Sum(l []int) int {
	val := 0
	for _, v := range l {
		val += v
	}
	return val
}

func countOccurrances(val int, l []int) int {
	i := 0
	for _, v := range l {
		if v == val {
			i += 1
		}
	}

	return i
}

func process2(l []int, r []int) int {
	sum := 0
	for _, v := range l {
		occurrances := countOccurrances(v, r)
		sum += v * occurrances

	}

	return sum
}

func process(l []int, r []int) int {
	l, lminVal := min(l)
	r, rminVal := min(r)
	val := lminVal - rminVal
	val = abs(val)

	if len(l) == 0 {
		return val
	}

	return val + process(l, r)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func min(l []int) ([]int, int) {
	minVal := l[0]
	minIdx := 0
	for idx, v := range l {
		if v < minVal {
			minVal = v
			minIdx = idx
		}
	}

	l = removeIndex(l, minIdx)

	return l, minVal
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
