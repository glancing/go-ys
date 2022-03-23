package loader

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var globalProxyArray []string

func ReturnLoadedProxies() []string {
	if globalProxyArray != nil {
		return globalProxyArray
	}
	proxies, err := ioutil.ReadFile("./proxies.txt")
	if err != nil {
		fmt.Println("err", err)
	}
	proxyArray := strings.Split(strings.ReplaceAll(string(proxies), "\r\n", "\n"), "\n")
	globalProxyArray = proxyArray
	return proxyArray
}