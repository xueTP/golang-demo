package main

import (
	"encoding/json"
	"fmt"
	"golang-demo/reptile"
)

func main() {
	// closeBag.FeibonaciNumsList()
	// errorDefer.HttpErrDemo()
	// maze.MazeBaseBreadthFirstDemo()
	reptile.GetHtml()
	//TestRange()
	//jsonDecode()
	//fmt.Printf("%v", uniquePathsWithObstacles([][]int{{0,0}, {1,0}}))
}

func TestRange() {
	a := []int8{1, 2, 3, 4, 5, 6}
	for i, v := range a {
		if v == 4 || v == 5 {
			a = append(a[:i], a[i+1:]...)
		}
	}
	fmt.Printf("%v", a)
}

func jsonDecode() {
	jsonString := `{"0":55}`
	priceJson := make(map[string]float64)
	json.Unmarshal([]byte(jsonString), &priceJson)
	fmt.Println(priceJson)
}

func uniquePaths(m int, n int) int {
	// reference https://segmentfault.com/a/1190000016315625?utm_medium=referral&utm_source=tuicool
	res := make([]int, m)
	res[0] = 1
	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			res[i] = res[i-1] + res[i]
		}
	}
	return res[m-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	res := make([]int, len(obstacleGrid[0]))
	res[0] = 1
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < n; i++ {
		if obstacleGrid[i][0] == 1 {
			res[0] = 0
		}
		for j := 1; j < m; j++ {
			res[j] = res[j-1] + res[j]
			if obstacleGrid[i][j] == 1 {
				res[j] = 0
			}
		}
	}
	return res[m-1]
}

// lc 75
func sortColors(nums []int) {
	res := make([]int, 3)
	for _, v := range nums {
		switch v {
		case 0:
			res[0]++
		case 1:
			res[1]++
		case 2:
			res[2]++
		}
	}
	index := 0
	for i, v := range res {
		for j := 1; j <= v; j++ {
			index += j
			nums[index] = i
		}
	}
}

// lc 77
func combine(n int, k int) [][]int {
	seed := make([]int, n)
	for i := 0; i < n; i++ {
		seed[i] = i + 1
	}
	res := [][]int{}
	combineRecursive(seed, []int{}, k, &res)
	return res
}

func combineRecursive(seed, item []int, t int, res *[][]int) {
	if t == 0 {
		*res = append(*res, item)
		return
	}
	for i, v := range seed {
		newSeed := append([]int{}, seed[i+1:]...)
		newItem := append([]int{}, item...)
		newItem = append(newItem, v)
		combineRecursive(newSeed, newItem, t-1, res)
	}
}
