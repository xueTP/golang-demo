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
		var temp int
		fmt.Fscanln(fileHandle, &temp)
		for j := range res[i] {
			// fmt.Fscanln(fileHandle, &res[i][j])
			fmt.Fscanf(fileHandle, "%d", &res[i][j])
		}
	}
	return res
}

type point struct {
	i, j int
}

func (p point) Add(p1 point) point {
	return point{i: p.i + p1.i, j: p.j + p1.j}
}

func checkPointLicit(p point, rowLimit, colLimit int) bool {
	if p.i < 0 || p.i >= rowLimit || p.j < 0 || p.j >= colLimit {
		return false
	}
	return true
}

func breadthFirstFindMazeUntie(mazeMap [][]int, start, end point, direction []point) [][]int {
	var rowLimit, colLimit int = len(mazeMap), len(mazeMap[0])
	if !checkPointLicit(start, rowLimit, colLimit) {
		panic("this start point is overflow")
	}
	if !checkPointLicit(end, rowLimit, colLimit) {
		panic("this end point is overflow")
	}

	res := make([][]int, len(mazeMap))
	for i := range res {
		res[i] = make([]int, len(mazeMap[0]))
	}
	q := []point{start}
	var curr point
	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		// res[curr.i][curr.j] = step

		for _, v := range direction {
			tempPoint := curr.Add(v)
			if !checkPointLicit(tempPoint, rowLimit, colLimit) {
				continue
			}
			if mazeMap[tempPoint.i][tempPoint.j] == 1 || res[tempPoint.i][tempPoint.j] > 0 {
				continue
			}
			if tempPoint == start {
				continue
			}
			q = append(q, tempPoint)
			res[tempPoint.i][tempPoint.j] = res[curr.i][curr.j] + 1
		}
	}
	return res
}

func MazeBaseBreadthFirstDemo() {
	mazeMap := readeMazeMap("maze/mazeMap")
	fmt.Println(mazeMap)
	direction := []point{{0, -1}, {i: 1, j: 0}, {i: 0, j: 1}, {i: -1, j: 0}}
	res := breadthFirstFindMazeUntie(mazeMap, point{i: 0, j: 0}, point{5, 4}, direction)
	fmt.Println(res)
}
