package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
	"github.com/DestroyerBots/httpclient/net/http"
	"github.com/glancing/go-ys/bot"
	"github.com/glancing/go-ys/reqClient"
	"github.com/glancing/go-ys/tasks/yeezysupply/types"
	"github.com/glancing/go-ys/tasks/yeezysupply/utils"
)

//if you wnat to call function again, just call like this wherever u need it
//loadYsConfig(task, internal)

func PushYeezySupplyHandlers(task *bot.Task) {
	task.Handlers["start"] = startYeezySupply
	task.Handlers["loadYsConfig"] = loadYsConfig
	task.Handlers["loadHome"] = loadHome
	task.Handlers["loadBloom"] = loadBloom
	task.Handlers["loadProduct"] = loadProduct
	task.Handlers["pollQueue"] = pollQueue
}

func startYeezySupply(task *bot.Task, internal *bot.TaskInternal) string {
	reqClient.SetTaskTransport(task.Proxy, task)
	client, err := reqClient.FetchClient(task, 30)
	if err != nil {
		fmt.Println("error creating task client")
		return "start"
	}
	task.Client = client
	parsedUrl, err := url.Parse("https://www.yeezysupply.com")
	if err != nil {
		fmt.Println("error parsing task url")
	}
	internal.ParsedUrl = parsedUrl
	fmt.Println("YS TASK STARTED", task.Id, task.Sku)
	return "loadYsConfig"
}

func loadYsConfig(task *bot.Task, internal *bot.TaskInternal) string {
	fmt.Println("loading ys config")

	headers := map[string][]string{
		"accept": {"application/json"},
	}

	reqUrl, _ := url.Parse("https://qzvvdaoys0.execute-api.us-east-2.amazonaws.com/prod/config/yeezysupply")
	defer task.Client.CloseIdleConnections()
	req := &http.Request{
		Method: "GET",
		URL: reqUrl,
		Header: headers,
		Close: true,
	}
	resp, err := task.Client.Do(req)
	if err != nil {
		fmt.Println(task.Id, "e", err)
		time.Sleep(5000 * time.Millisecond)
		return "loadConfig"
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("e", err)
		time.Sleep(5000 * time.Millisecond)
		return "loadConfig"
	}

	respBody := string(bodyBytes)

	var parsedConfig types.ApiYsConfig
	json.Unmarshal([]byte(respBody), &parsedConfig)
	internal.Config = types.ApiYsConfig(parsedConfig)	

	fmt.Println("loaded config")

	return "loadHome"
}

func loadHome(task *bot.Task, internal *bot.TaskInternal) string {
	fmt.Println("loading home")

	headers := map[string][]string{
    "sec-ch-ua": {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\""},
    "sec-ch-ua-mobile": {"?0"},
    "sec-ch-ua-platform": {"\"Windows\""},
    "upgrade-insecure-requests": {"1"},
    "user-agent": {internal.Config.UserAgentArray[0]},
    "accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
    "sec-fetch-site": {"none"},
    "sec-fetch-mode": {"navigate"},
    "sec-fetch-user": {"?1"},
    "sec-fetch-dest": {"document"},
    "accept-encoding": {"gzip, deflate, br"},
    "accept-language": {"en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"},
		"HEADERORDER": {"sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "upgrade-insecure-requests", "user-agent", "accept", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-user", "sec-fetch-dest", "accept-encoding", "accept-language", "Cookie"},
		"PSEUDOORDER": {":method", ":authority", ":scheme", ":path"},
	}

	reqUrl, _ := url.Parse("https://www.yeezysupply.com")
	req := &http.Request{
		Method: "GET",
		URL: reqUrl,
		Header: headers,
	}
	resp, err := task.Client.Do(req)
	if err != nil {
		fmt.Println("e", err)
		time.Sleep(5000 * time.Millisecond)
		return "loadHome"
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error loading home", resp.StatusCode)
		time.Sleep(5000 * time.Millisecond)
		return "loadHome"
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("e", err)
		time.Sleep(5000 * time.Millisecond)
		return "finished"
	}

	respBody := string(bodyBytes)
	if (strings.Contains(respBody, "UNFORTUNATELY WE ARE UNABLE TO GIVE YOU ACCESS TO OUR SITE AT THIS TIME")) {
		fmt.Println("task banned (home), restarting")
		return "finished"
	}

	pixelData := utils.ParsePixel(respBody)
	if (!pixelData.PixelFound) {
		fmt.Println("no pixel found")
	} else {
		fmt.Println("pixel data found on home")
		internal.PixelConfig = pixelData
	}

	return "loadBloom"
}

func loadBloom(task *bot.Task, internal *bot.TaskInternal) string {
	fmt.Println("loading bloom")

	headers := map[string][]string{
		"x-instana-t": {"f38b101ee67b5b84"},
		"sec-ch-ua-mobile": {"?0"},
		"user-agent": {internal.Config.UserAgentArray[0]},
		"x-instana-l": {"1,correlationType=web;correlationId=f38b101ee67b5b84"},
		"x-instana-s": {"f38b101ee67b5b84"},
		"content-type": {"application/json"},
		"sec-ch-ua-platform": {"\"macOS\""},
		"sec-ch-ua": {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\""},
		"accept": {"*/*"},
		"sec-fetch-site": {"same-origin"},
		"sec-fetch-mode": {"cors"},
		"sec-fetch-dest": {"empty"},
		"referer": {"https://www.yeezysupply.com/"},
		"accept-encoding": {"gzip, deflate, br"},
		"accept-language": {"en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"},
		"HEADERORDER": {"x-instana-t", "sec-ch-ua-mobile", "user-agent", "x-instana-l", "x-instana-s", "content-type", "sec-ch-ua-platform", "sec-ch-ua", "accept", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "referer", "accept-encoding", "accept-language", "Cookie"},
		"PSEUDOORDER": {":method", ":authority", ":scheme", ":path"},
	}

	reqUrl, _ := url.Parse("https://www.yeezysupply.com/api/yeezysupply/products/bloom")
	req := &http.Request{
		Method: "GET",
		URL: reqUrl,
		Header: headers,
	}
	resp, err := task.Client.Do(req)
	if err != nil {
		fmt.Println("error loading bloam")
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error loading bloom", resp.StatusCode)
		time.Sleep(5000 * time.Millisecond)
		return "loadBloom"
	}

	defer resp.Body.Close()

	return "loadProduct"
}

func loadProduct(task *bot.Task, internal *bot.TaskInternal) string {
	fmt.Println("loading product")

	headers := map[string][]string{
    "sec-ch-ua": {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\""},
    "sec-ch-ua-mobile": {"?0"},
    "sec-ch-ua-platform": {"\"Windows\""},
    "upgrade-insecure-requests": {"1"},
    "user-agent": {internal.Config.UserAgentArray[0]},
    "accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
    "sec-fetch-site": {"none"},
    "sec-fetch-mode": {"navigate"},
    "sec-fetch-user": {"?1"},
    "sec-fetch-dest": {"document"},
    "accept-encoding": {"gzip, deflate, br"},
    "accept-language": {"en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"},
		"HEADERORDER": {"sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "upgrade-insecure-requests", "user-agent", "accept", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-user", "sec-fetch-dest", "accept-encoding", "accept-language", "Cookie"},
		"PSEUDOORDER": {":method", ":authority", ":scheme", ":path"},
	}
	reqUrl, _ := url.Parse(fmt.Sprintf("https://www.yeezysupply.com/product/%s", task.Sku))
	req := &http.Request{
		Method: "GET",
		URL: reqUrl,
		Header: headers,
	}
	resp, err := task.Client.Do(req)
	if err != nil {
		fmt.Println("e", err)
		return "finished"
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error loading product", resp.StatusCode)
		time.Sleep(5000 * time.Millisecond)
		return "finished"
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("e", err)
		return "finished"
	}

	respBody := string(bodyBytes)

	if (strings.Contains(respBody, "UNFORTUNATELY WE ARE UNABLE TO GIVE YOU ACCESS TO OUR SITE AT THIS TIME")) {
		fmt.Println("task banned (product), restarting")
		time.Sleep(5000 * time.Millisecond)
		return "finished"
	}

	if (!internal.PixelConfig.PixelFound) {
		pixelData := utils.ParsePixel(respBody)
		if (!pixelData.PixelFound) {
			fmt.Println("restarting task")
			time.Sleep(3000 * time.Millisecond)
			return "start"
		} else {
			fmt.Println("found pixel on product")
			internal.PixelConfig = pixelData
		}
	}
	return "pollQueue"
}

func pollQueue(task *bot.Task, internal *bot.TaskInternal) string {
	fmt.Println("polling queue")

	headers := map[string][]string{
		"sec-ch-ua": {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\""},
		"accept": {"application/json, text/plain, */*"},
		"sec-ch-ua-mobile": {"?0"},
    "user-agent": {internal.Config.UserAgentArray[0]},
		"sec-ch-ua-platform": {"\"macOS\""},
		"sec-fetch-site": {"same-origin"},
		"sec-fetch-mode": {"cors"},
		"sec-fetch-dest": {"empty"},
		"referer": {fmt.Sprintf("https://www.yeezysupply.com/product/%s", task.Sku)},
		"accept-encoding": {"gzip, deflate, br"},
		"accept-language": {"en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7"},
		"HEADERORDER": {"sec-ch-ua", "accept", "sec-ch-ua-mobile", "user-agent", "sec-ch-ua-platform", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "referer", "accept-encoding", "accept-language", "Cookie"},
		"PSEUDOORDER": {":method", ":authority", ":scheme", ":path"},
	}
	reqUrl, _ := url.Parse("https://www.yeezysupply.com/__queue/yzysply")
	req := &http.Request{
		Method: "GET",
		URL: reqUrl,
		Header: headers,
	}
	resp, err := task.Client.Do(req)
	if err != nil {
		fmt.Println("e", err)
		return "finished"
	}

	defer resp.Body.Close()

	passed := checkForPassedCookie(task.Client.Jar.Cookies(internal.ParsedUrl))
	if !passed {
		time.Sleep(3000 * time.Millisecond)
		return "pollQueue"
	}

	return "finished"
}

func checkForPassedCookie(cookies []*http.Cookie) bool {
	for _, c := range cookies {
		if strings.Contains(c.Name, "_u") {
			return true
		}
	}
	return false
}