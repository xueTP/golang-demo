package zhenai

import (
	"fmt"
	"golang-demo/reptile/engine"
	"regexp"
)

const cityRule = `<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]*)</a>`

func CityParser(body []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityRule)
	res := reg.FindAllSubmatch(body, -1)
	parserResult := engine.ParseResult{}
	for _, v := range res {
		fmt.Printf("ctiy: %s url: %s\n", v[2], v[1])
		request := engine.Request{Url: string(v[1]), ParserFunc: nil}
		parserResult.Requests = append(parserResult.Requests, request)
		parserResult.Item = append(parserResult.Item, string(v[2]))
	}
	return parserResult
}
