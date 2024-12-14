package src

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DayOne struct {
	FilePath string
}

func (d *DayOne) CalculateTotalDistance() int {
	leftList, rightList := d.buildLists()

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0

	// assuming both lists are the same length
	for i := 0; i < len(leftList); i++ {
		totalDistance += d.abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func (d *DayOne) CalculateTotalSimilarityScore() int {
	leftList, rightList := d.buildLists()
	occurrences := make(map[int][]int)

	// occurrences of the same value in leftList
	for _, val := range leftList {
		if _, ok := occurrences[val]; ok {
			occurrences[val][0]++
			continue
		}
		occurrences[val] = []int{1, 0}
	}

	// count occurrences in in right list
	for _, val := range rightList {
		if _, ok := occurrences[val]; ok {
			occurrences[val][1]++
		}
	}

	total := 0
	for key, val := range occurrences {
		if val[1] > 0 {
			total += (val[0] * (key * val[1]))
		}
	}

	return total
}

func (d *DayOne) openFile() *os.File {
	file, err := os.Open(d.FilePath)
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	return file
}

func (d *DayOne) buildLists() ([]int, []int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	file := d.openFile()

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineIds := strings.Split(scanner.Text(), "   ")
		leftVal, _ := strconv.Atoi(lineIds[0])
		rightVal, _ := strconv.Atoi(lineIds[1])
		leftList = append(leftList, leftVal)
		rightList = append(rightList, rightVal)
	}
	return leftList, rightList
}

func (d *DayOne) abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
