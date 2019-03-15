package reptile

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang-demo/reptile/engine"
	"golang-demo/reptile/util"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
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
	// m := engine.Request{Url: "http://city.zhenai.com/", ParserFunc: zhenai.CityListParser}
	// engine.ConcurrentQueueEngine{WorkCount: 10, Scheduling: engine.Scheduling{}}.Run(m)
	//engine.Run(m)
	http.HandleFunc("/", rootFunc)
	http.ListenAndServe(":8888", nil)
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./reptile/view/html/showlist.html"))
	//body, _ := util.Fetch("http://album.zhenai.com/u/108757455")
	//res := zhenai.PersonParser(body, "Abigale", "http://album.zhenai.com/u/108757455")
	getQuery := r.URL.Query()
	var search, from string
	if v, ok := getQuery["search"]; ok {
		search = v[0]
	}
	if v, ok := getQuery["from"]; ok {
		from = v[0]
	}
	fmt.Println(search, from)
	offset, _ := strconv.Atoi(from)
	if offset < 0 {
		offset = 0
	}
	res, total := engine.GetList(search, offset, 10)
	fmt.Printf("res: %v，\n totle: %d \n, search: %v \n", res, total, search)
	data := map[string]interface{}{
		"list":    res,
		"count":   total,
		"search":  search,
		"from":    offset,
		"pageUrl": "http://127.0.0.1:8888/?search=" + search + "&from=",
		"prev":    offset - 10,
		"curr":    offset + 10,
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
