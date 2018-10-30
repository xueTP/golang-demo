package closeBag

import (
	"fmt"
	"io"
	"bufio"
	"strconv"
	"strings"
)

//斐波那契数列
//  1， 1， 2， 3， 5 ...  n = n-1 + n-2

type intGen func() int

func (ig intGen) Read(p []byte) (n int, err error) {
	next := ig()
	if next > 10000 {
		return 0, io.EOF
	}
	s := strconv.Itoa(next) + "\n"
	return strings.NewReader(s).Read(p)
}

func feiboaci() intGen {
	var a, b int = 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func PrintFailContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func FeibonaciNumsList() {
	f := feiboaci()
	PrintFailContents(f)
}
