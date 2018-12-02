package util

import (
	"bufio"
	"errors"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

// determineEncoding 猜测reader(html)使用的编码，返回相应的解码->utf8
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	buff, err := reader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(buff, "")
	return e
}

// Fetch 通过url获取页面的内容
func Fetch(url string) ([]byte, error) {
	resp := MockBrowser(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("this http the statusCode is %v", resp.StatusCode)
		return nil, errors.New("this html fetch fail")
	}
	bufioReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(bufioReader).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func MockBrowser(url string) *http.Response {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return response
}
