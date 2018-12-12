package main

import (
	"fmt"
	"encoding/json"
	"golang-demo/reptile"
)

func main() {
	// closeBag.FeibonaciNumsList()
	// errorDefer.HttpErrDemo()
	// maze.MazeBaseBreadthFirstDemo()
	reptile.GetHtml()
	//TestRange()
	//jsonDecode()
}

func TestRange() {
	a := []int8{1,2,3,4,5,6}
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
