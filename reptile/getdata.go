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
	"fmt"
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
    logrus.Info("tttt")
	t := template.Must(template.ParseFiles("./reptile/view/html/showlist.html"))
	//body, _ := util.Fetch("http://album.zhenai.com/u/108757455")
	//res := zhenai.PersonParser(body, "Abigale", "http://album.zhenai.com/u/108757455")
	res, totle := engine.GetList("", 0)
	fmt.Printf("res: %v，\n totle: %d \n", res, totle)
	data := map[string]interface{}{
		"list": res,
		"count": totle,
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
