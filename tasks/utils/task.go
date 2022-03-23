package utils

import (
	"github.com/glancing/go-ys/loader"
	"sync"
	"fmt"
	"strings"
	"github.com/glancing/go-ys/bot"
)

var rotateAccess sync.Mutex

func SelectProxy(proxiesLoaded []string) string {
	proxyString := PickRandomIndex(proxiesLoaded)
	splitProxy := strings.Split(proxyString, ":")
	var formattedProxy string
	switch(len(splitProxy)) {
		case 2:
			formattedProxy = fmt.Sprintf("http://%s:%s", splitProxy[0], splitProxy[1])
		case 4:
			formattedProxy = fmt.Sprintf("http://%s:%s@%s:%s", splitProxy[2], splitProxy[3], splitProxy[0], splitProxy[1])
		default:
			formattedProxy = ""
	}
	return formattedProxy
}

func RotateProxy(task *bot.Task) {
	rotateAccess.Lock()
	defer rotateAccess.Unlock()
	proxies := loader.ReturnLoadedProxies()
	proxy := SelectProxy(proxies)
	task.Proxy = proxy
}