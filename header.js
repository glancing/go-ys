let headerString = `
"x-instana-t": {"f38b101ee67b5b84"},
"sec-ch-ua-mobile": {"?0"},
"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36"},
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
`

headerString = headerString.replace(/\s+/g, '');
let headers = headerString.split(',"');
let keys = headers.map(header => {
  return header.split(":")[0].replace(/["']/g, '');
});
keys = keys.map(key => {
  return `"` + key + `"`
});
console.log(keys.join(", "))