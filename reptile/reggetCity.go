package reptile

import (
	"regexp"
	"fmt"
)

func regGetCity(buff []byte) {
	reg := regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]*)</div>`)
	res := reg.FindAllSubmatch(buff, -1)
	fmt.Printf("%+s", res)
	// for _, v := range res {
	// 	fmt.Printf("ctiy: %s url: %s\n", v[2], v[1])
	// }
	// zhenai.PersonParser(buff)
}
