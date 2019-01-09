package reptile

import (
	"bufio"
	"github.com/Sirupsen/logrus"
	"golang-demo/reptile/util"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"html/template"
	"golang-demo/reptile/engine"
	"golang-demo/reptile/parser/zhenai"
)

func getHtml(url string) {
	resp := util.MockBrowser(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("this http the statusCode is %v", resp.StatusCode)
		return
	}
	bufioReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, determineEncoding(bufioReader).NewDecoder())
	body, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s", body)
	regGetCity(body)
}

func GetHtml() {
	// getHtml("http://album.zhenai.com/u/1571428123")
	m := engine.Request{Url: "http://city.zhenai.com/", ParserFunc: zhenai.CityListParser}
	engine.ConcurrentQueueEngine{WorkCount: 10, Scheduling: engine.Scheduling{}}.Run(m)
	//engine.Run(m)
	http.HandleFunc("/", rootFunc)
	http.ListenAndServe(":8888", nil)
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./reptile/view/html/showlist.html"))
	data := map[string]interface{}{
		"list": []int{1,2,3,4},
	}
	t.Execute(w, data)
}

func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	buff, err := reader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(buff, "")
	return e
}
