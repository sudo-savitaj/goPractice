package main

import . "fmt"

const INT_MAX = 10000000

func main() {
	var squareSize int
	var timeLimit int
	var cordsTime = [][]int{}

	Scanf("%d", &squareSize)
	Scanf("%d", &timeLimit)

	for i := 0; i < squareSize; i++ {
		row := make([]int, squareSize) // create a new slice to add next row values
		for j, _ := range row {
			Scanf("%d", &row[j])
		}
		cordsTime = append(cordsTime, row)
	}
	graph := buildGraph(cordsTime,squareSize)

}
func buildGraph(cordsTime [][]int, size int) [][]int {
	var graph [][]int
	for x:=0; x<size; x++ {
		for y:=0; y<size; y++ {
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

			for j := endYCord; j >= startYCord; j-- {
				for i := endXCord; i >= startXCord; i-- {
					graph[]
				}
			}
		}
	}

	return graph
}

func minDistance(dist []int, size int, sptSet []bool) int {
	min := INT_MAX;
	var min_index int;
	for v := 0; v < size; v++ {
		if (sptSet[v] == false && dist[v] <= min) {
			min = dist[v];
			min_index = v;
		}
	}
	return min_index;
}

func printSolution(dist []int, n int) {
	Println("Vertex   Distance from Source\n");
	for i := 0; i < n; i++ {
		Println("%d tt %d\n", i, dist[i]);
	}
}

func dijkstra(graph [][]int, src int, size int) {
	var dist []int; // The output array.  dist[i] will hold the shortest
	// distance from src to i

	var sptSet []bool; // sptSet[i] will be true if vertex i is included in shortest
	// path tree or shortest distance from src to i is finalized

	// Initialize all distances as INFINITE and stpSet[] as false
	for i := 0; i < size; i++ {
		dist[i] = INT_MAX;
		sptSet[i] = false;
	}
	// Distance of source vertex from itself is always 0
	dist[src] = 0;

	// Find shortest path for all vertices
	for count := 0; count < size-1; count++ {
		// Pick the minimum distance vertex from the set of vertices not
		// yet processed. u is always equal to src in the first iteration.
		u := minDistance(dist, size, sptSet);

		// Mark the picked vertex as processed
		sptSet[u] = true;

		// Update dist value of the adjacent vertices of the picked vertex.
		for v := 0; v < size; v++ {

			// Update dist[v] only if is not in sptSet, there is an edge from
			// u to v, and total weight of path from src to  v through u is
			// smaller than current value of dist[v]
			if (!sptSet[v] && graph[u][v] != 0) && dist[u] != INT_MAX && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v];
			}
		}

	}
	// print the constructed distance array
	printSolution(dist, size);
}
