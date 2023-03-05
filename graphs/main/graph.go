package main

import (
	"fmt"
	"main/graphs"
)

func main() {

	var ans int = graphs.NumIslands([][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}})
	fmt.Printf("Number of Islands: %d", ans)

}
