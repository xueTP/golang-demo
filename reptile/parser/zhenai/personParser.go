package zhenai

import (
	"golang-demo/reptile/engine"
	"regexp"
	"golang-demo/reptile/parser"
	"golang-demo/reptile/model"
	"strconv"
	"github.com/Sirupsen/logrus"
)

const basePersonRule = `<div class="m-btn purple" data-v-ff544c08>([^<]*)</div>`
const basePersonOtherRule = `<div class="m-btn pink" data-v-ff544c08>([^<]*)</div>`

func PersonParser(body []byte) engine.ParseResult {
	intReg := regexp.MustCompile(`([0-9]*)`)
	baseReg := regexp.MustCompile(basePersonRule)
	resBase := baseReg.FindAllSubmatch(body, -1)
	logrus.Errorf("resbase : %v", resBase)
	parson := model.Person{}
	parserResult := engine.ParseResult{}
	var baseString []byte
	for i, v := range resBase {
		switch i {
		case 0:
			if string(v[1]) == "未婚" {
				parson.IsMarred = false
			} else {
				parson.IsMarred = true
			}
		case 1:
			ageRes := intReg.Find(v[1])
			parson.Age, _ = strconv.Atoi(string(ageRes))
		case 2:
			parson.Constellation = string(v[1])
		case 3:
			heightRes := intReg.Find(v[1])
			parson.Height, _ = strconv.Atoi(string(heightRes))
		case 4:
			weightRes := intReg.Find(v[1])
			parson.Weight, _ = strconv.Atoi(string(weightRes))
		case 5:
			parson.WorkAddress = string(v[1])
		case 6:
			parson.Income = string(v[1])
		case 7:
			parson.Job = string(v[1])
		case 8:
			parson.Education = string(v[1])
		}
		v[1] = append(v[1], []byte{','}...)
		baseString = append(baseString, v[1]...)
	}
	otherReg := regexp.MustCompile(basePersonOtherRule)
	resOther := otherReg.FindAllSubmatch(body, -1)
	var baseOther []byte
	for _, v := range resOther {
		v[1] = append(v[1], []byte{','}...)
		baseOther = append(baseOther, v[1]...)
	}
	parson.Content = string(baseString)
	parson.OtherDetail = string(baseOther)
	logrus.Infof("parson: %+v \n", parson)
	request := engine.Request{ParserFunc: parser.NilParserFunc}
	parserResult.Requests = append(parserResult.Requests, request)
	parserResult.Item = append(parserResult.Item, parson)
	return parserResult
}
