package reptile

import (
	"fmt"
	"regexp"
)

func regGetCity(buff []byte) {
	reg := regexp.MustCompile(`<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]*)</a>`)
	res := reg.FindAllSubmatch(buff, -1)
	for _, v := range res {
		fmt.Printf("ctiy: %s url: %s\n", v[2], v[1])
	}
}
