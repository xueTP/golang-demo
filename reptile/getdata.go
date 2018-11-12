package reptile

import (
	"net/http"
	"io/ioutil"
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
	e, reader := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s", body)
	regGetCity(body)
}

func GetHtml() {
	getHtml("http://city.zhenai.com/")
}

func determineEncoding(reader io.Reader) (encoding.Encoding, io.Reader) {
	r := bufio.NewReader(reader)
	buff, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(buff, "")
	return e, r
}