package reptile

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
)

func getHtml(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("this http the statusCode is %v", resp.StatusCode)
		return
	}
	utf8Reader := transform.NewReader(resp.Body, Determineencoding(resp.Body).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}

func GetHtml() {
	getHtml("http://xm.zu.fang.com")
}

func Determineencoding(reader io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}