package main

import "fmt"

func main() {
	var squareSize int
	var timeLimit int
	var cordsTime = [][]int{}

	fmt.Scanf("%d", &squareSize)
	fmt.Scanf("%d", &timeLimit)

	for i := 0; i < squareSize; i++ {
		row := make([]int, squareSize) // create a new slice to add next row values
		for j, _ := range row {
			fmt.Scanf("%d", &row[j])
		}
		cordsTime = append(cordsTime, row)
	}

	maxXCord, maxYCord := exploreHackerLand(squareSize, timeLimit, cordsTime)
	fmt.Println(maxXCord, maxYCord)
	maximumArea := (maxXCord + 1) * (maxYCord + 1)

	fmt.Println(maximumArea)
}

func exploreHackerLand(size int, timeLimit int, cordsTime [][]int) (int, int) {
	var visitedCords = [][]int{}
	timeGain := 0
	maxXCord := 0
	maxYCord := 0
	maxArea := (maxXCord + 1 ) * (maxYCord + 1)
	x := 0
	y := 0

	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			row[j] = 0
		}
		visitedCords = append(visitedCords, row)
	}

	for timeLimit > timeGain {
		x, y = nextMinimumCord(cordsTime, visitedCords, size, x, y, maxArea)
		if (timeGain+cordsTime[x][y] > timeLimit) {
			break
		}

		timeGain = timeGain + cordsTime[x][y]
		maxXCord = max(maxXCord, x)
		maxYCord = max(maxYCord, y)
		visitedCords[x][y] = 1
	}
	return maxXCord, maxYCord
}
func nextMinimumCord(cordsTime [][]int, visitedCords [][]int, size int, x int, y int, maxArea int) (int, int) {
	startXCord := x - 1
	if startXCord < 0 {
		startXCord = 0
	}
	endXCord := x + 1
	if endXCord >= size {
		endXCord = size - 1
	}

	startYCord := y - 1
	if startYCord < 0 {
		startYCord = 0
	}
	endYCord := y + 1
	if endYCord >= size {
		endYCord = size - 1
	}
	minXCord := 0
	minYCord := 0
	minTime := 10000

		for j := endYCord; j >= startYCord; j-- {
			for i := endXCord; i >= startXCord; i-- {
				tempArea := (x + 1) * (y + 1)
				if visitedCords[i][j] == 0 && cordsTime[i][j] < minTime && (x != i || y != j) && tempArea > maxArea {
					minXCord = i
					minYCord = j
					minTime = cordsTime[i][j]
				}
			}
		}

	return minXCord, minYCord
}

func max(value1 int, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}
