package util

import (
	"bufio"
	"errors"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

// determineEncoding 猜测reader(html)使用的编码，返回相应的解码->utf8
func determineEncoding(reader io.Reader) (encoding.Encoding, io.Reader) {
	r := bufio.NewReader(reader)
	buff, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(buff, "")
	return e, r
}

// Fetch 通过url获取页面的内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("this http the statusCode is %v", resp.StatusCode)
		return nil, errors.New("this html fetch fail")
	}
	e, reader := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}
