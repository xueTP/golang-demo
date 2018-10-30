package closeBag

import "fmt"

// 通过设置函数外部的自由变量sum 来保存每次累加的结果
// sum被函数内部引用
func adder() func (int) int {
	var sum int = 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func SumByCloseBag() {
	add := adder()
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(add(i))
	}
}

// 函数式编程（严格没有变量/赋值等操作）
type iAdder func (int) (int, iAdder)

func adderFunc(v int) iAdder {
	return func(i int) (int, iAdder) {
		return v + i, adderFunc(i + v )
	}
}

func SumStrictFuncStyle() {
	a := adderFunc(0)
	var val int
	for i := 0 ; i < 10 ; i ++ {
		val, a = a(i)
		fmt.Println(val)
	}
}