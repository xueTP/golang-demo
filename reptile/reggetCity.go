package reptile

import "golang-demo/reptile/parser/zhenai"

func regGetCity(buff []byte) {
	// reg := regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([^<]*)</div>`)
	// res := reg.FindAllSubmatch(buff, -1)
	// fmt.Printf("%+s", res)
	// for _, v := range res {
	// 	fmt.Printf("ctiy: %s url: %s\n", v[2], v[1])
	// }
	zhenai.PersonParser(buff)
}
