package loader

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func ReturnLoadedProxies() []string {
	proxies, err := ioutil.ReadFile("./proxies.txt")
	if err != nil {
		fmt.Println(err)
	}
	proxyArray := strings.Split(string(proxies), "\n")
	return proxyArray
}