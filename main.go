package main

import (
	"bou.ke/monkey"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	// closeBag.FeibonaciNumsList()
	// errorDefer.HttpErrDemo()
	// maze.MazeBaseBreadthFirstDemo()
	//reptile.GetHtml()
	//TestRange()
	//jsonDecode()
	//fmt.Printf("%v", uniquePathsWithObstacles([][]int{{0,0}, {1,0}}))
	//jsonRpc.JsonRpcMain()
	//fmtScan()
	//bitOperation()
	monkey_test()
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

func fmtScan() { // 类似c的cin <<
	var name string
	var age int
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)
	fmt.Printf("%s的年龄是%d", name, age)
}

func bitOperation() {
	var a = 1 >> 2
	var b = -1 >> 2
	var c = 1 << 2
	var d = -1 << 2
	fmt.Println(a, b, c, d)
	fmt.Printf("%b, %b, %b, %b", a, b, c, d)
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

// lc 87
func merge(nums1 []int, m int, nums2 []int, n int) {
	k, i, j := m+n-1, m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if i < 0 && j >= 0 {
		for j >= 0 {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
}

// lc 200
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	num, il, jl := 0, len(grid), len(grid[0])
	dire := [][]int{{0, -1}, {1,0}, {0, 1}, {-1, 0}}
	for i, line := range grid {
		for j, v := range line {
			if v == '1' {
				grid[i][j] = '0'
				num ++
				q := [][]int{{i, j}}
				for len(q) > 0 {
					now := q[0]
					q = q[1:]
					for _, val := range dire {
						if now[0] + val[0] >= 0 && now[0] + val[0] < il && now[1] + val[1] >= 0 && now[1] + val[1] < jl && grid[now[0] + val[0]][now[1] + val[1]] == '1' {
							grid[now[0] + val[0]][now[1] + val[1]] = '0'
							q = append(q, []int{now[0] + val[0], now[1] + val[1]})
						}
					}
				}
			}
		}
	}
	return num
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// lc 106
func buildTree(inorder []int, postorder []int) *TreeNode {
	return recursive106(inorder, postorder)
}

func recursive106(in, post []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}
	//最后一个
	l := len(post) - 1
	t := post[l]
	node := &TreeNode{Val: t}
	//in 里获取
	var n int
	for i, v := range in {
		if v == t {
			n = i
			break
		}
	}
	c := l - 1 - (l - 1 - n)
	node.Left = recursive106(in[:n], post[:c])
	node.Right = recursive106(in[n+1:], post[c:l])
	return node
}

type tt struct {
	dn []string
}

func t1() {
	tip := make(map[string]tt)
	tip["a"] = tt{dn: []string{"a"}}
	tip["b"] = tt{dn: []string{"b"}}
	for k, _ := range tip {
		node := tip[k]
		node.dn = append(node.dn, "e")
	}
	fmt.Printf("%+v", tip)
}

func monkey_test() {
	rnd, _ := rand.Int(rand.Reader, big.NewInt(100))
	val := int8(rnd.Int64())
	monkey.Patch(rnd.Int64, func() int64 {return int64(23232)})
	fmt.Print(val)
}