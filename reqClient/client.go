package reqClient

import (
	"net/url"
	"fmt"
	"time"
	"strings"
	"sync"
	http "github.com/DestroyerBots/httpclient/net/http"
	"github.com/glancing/go-ys/bot"
	"github.com/DestroyerBots/httpclient/net/http/cookiejar"
	tls "github.com/DestroyerBots/utls"
)

type CachedTransport struct {
	Transport *http.Transport
	Proxy string
}

var cachedTransports = []CachedTransport{}
var cacheAccess sync.Mutex

func findExistingTransport(proxyString string) *http.Transport {
	for _, c := range cachedTransports {
		if strings.EqualFold(c.Proxy, proxyString) {
			return c.Transport
		}
	}
	return nil
}


func SetTaskTransport(proxyString string, task *bot.Task) {
	cacheAccess.Lock()
	defer cacheAccess.Unlock()
	existingTransport := findExistingTransport(proxyString)
	if existingTransport == nil {
		fmt.Println("creating new transport")
		var clientProxy *url.URL
		if proxyString != "" {
			proxy, err := url.Parse(proxyString)
			if err != nil {
				fmt.Println("err parsing proxy", err)
			}
			fmt.Println("using proxy")
			clientProxy = proxy
		}
		transport := &http.Transport {
			ClientHelloID: &tls.HelloChrome_92,
			DisableCompression: false,
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{},
		}
	
		if clientProxy != nil {
			fmt.Println("setting proxy")
			transport.Proxy = http.ProxyURL(clientProxy)
		}
		cachedTransports = append(cachedTransports, CachedTransport{
			Transport: transport,
			Proxy: proxyString,
		})
		task.Transport = transport
	} else {
		task.Transport = existingTransport
	}
}


func FetchClient(task *bot.Task, timeout int) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err !=  nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Jar: jar,
		Transport: task.Transport,
	}
	return client, nil
}
