package maze

import (
	"os"
	"fmt"
)

func readeMazeMap(fileName string) [][]int {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	var row, col int
	fmt.Fscanf(fileHandle, "%d %d", &row, &col)
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
		for j := range res[i] {
			fmt.Fscanf(fileHandle, "%d", &res[i][j])
		}
	}
	return res
}

func MazeBaseBreadthFirstDemo() {
	mazeMap := readeMazeMap("maze/mazeMap")
	fmt.Println(mazeMap)
}
