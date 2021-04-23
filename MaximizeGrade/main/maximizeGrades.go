package main

import "fmt"

func main() {
	var n int
	var k int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	grades := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &grades[i])
	}
	middle := n - k
	start := middle / 2
	if (middle % 2) == 0{
		start --
	}
	end := (n - (middle / 2)) -1
	var maxGrade = -1
	for i := start; i <= end; i++ {
		maxGrade = max(maxGrade,grades[i])
	}
	fmt.Println(maxGrade)
}
func max(value1 int, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}












