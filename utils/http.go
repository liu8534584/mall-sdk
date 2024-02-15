package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpClient struct {
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

func GetResponse2(urls string, headers map[string]string, timeout time.Duration, maxRetry int) ([]byte, error) {
	i := 0
	for {
		if i > maxRetry {
			return nil, errors.New("请求失败")
		}
		request, _ := http.NewRequest("GET", urls, nil)
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Connection", "keep-alive")
		if len(headers) > 0 {
			for k, v := range headers {
				request.Header.Set(k, v)
			}
		}
		//设置超时时间
		client := &http.Client{
			Timeout: timeout,
		}
		response, err := client.Do(request)
		if response != nil {
			defer response.Body.Close()
		}
		if err != nil || response.StatusCode != 200 {
			i++
			continue
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			i++
			continue
		}
		return body, nil
	}
}

func GetResponse(urls string, headers map[string]string, timeout time.Duration) ([]byte, error) {
	maxRetry := 3
	i := 0
	for {
		if i > maxRetry {
			return nil, errors.New("请求失败")
		}

		request, _ := http.NewRequest("GET", urls, nil)
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Set("Connection", "keep-alive")
		if len(headers) > 0 {
			for k, v := range headers {
				request.Header.Set(k, v)
			}
		}
		//设置超时时间
		client := &http.Client{
			Timeout: timeout,
		}
		response, err := client.Do(request)
		if response != nil {
			defer response.Body.Close()
		}
		if err != nil || response.StatusCode != 200 {
			i++
			continue
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			i++
			continue
		}
		return body, nil
	}

}

func HttpGet(urls string, timeout time.Duration) (string, error) {
	req, err := http.NewRequest(http.MethodGet, urls, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{Timeout: timeout}
	do, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HttpPost(urls string, data map[string]string) (string, error) {
	var c *http.Client = &http.Client{

		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost: 10,
			//测试 暂时修改超时时间
			ResponseHeaderTimeout: time.Second * 5,
		},
	}

	var clusterInfo = url.Values{}
	for k, v := range data {
		clusterInfo.Add(k, v)
	}
	params := clusterInfo.Encode()
	resp, err := c.Post(urls, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPostJson(url string, data string, timeout time.Duration, headerMap map[string]string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	//设置请求头
	if len(headerMap) == 0 {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		for k, v := range headerMap {
			req.Header.Set(k, v)
		}
	}

	//发送请求
	resp, err := client.Do(req)
	if resp == nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Dingding(url string, msg string) error {
	message := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	_, err := HttpPostJson(url, message, time.Second*3, nil)
	if err != nil {
		return err
	}
	return nil
}

func WechatAlarmMsg(key string, msg string) error {
	message := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	_, err := HttpPostJson("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+key, message, time.Second*3, nil)
	if err != nil {
		return err
	}
	return nil
}

// 获得html内容
func GetHtml(url string, headers map[string]string) ([]byte, error) {
	var logInfo string
	response, err := GetResponse(url, headers, 3*time.Second)
	if err != nil {
		logInfo = fmt.Sprintf("http 请求失败，url:%v,err:%v", url, err)
		return nil, errors.New(logInfo)
	}

	return response, nil
}

func MakeHeaders(headers string) map[string]string {
	mm := make(map[string]string)
	if len(headers) == 0 {
		return mm
	}
	m := strings.Split(headers, "\n")
	for _, v := range m {
		kk := strings.Split(v, ":")
		if len(kk) > 1 {
			mm[kk[0]] = mm[kk[1]]
		}
	}
	return mm
}

func SetQuery(link string, params map[string]string) string {
	if params == nil {
		return link
	}

	querys := getQuerys(link)

	for k, v := range params {
		querys[k] = v
	}

	var newParam []string
	for k, v := range querys {
		newParam = append(newParam, k+"="+url.QueryEscape(v))
	}

	param := strings.Join(newParam, "&")

	if strings.Index(link, "?") == -1 {
		link += "?" + param
	} else {
		split := strings.Split(link, "?")
		link = split[0] + "?" + param
	}
	return link

}

func GetQuery(link string, key string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}
	if m[key] == nil {
		return ""
	}

	return m[key][0]
}

func getQuerys(link string) map[string]string {
	u, err := url.Parse(link)
	var mm = make(map[string]string)
	if err != nil {
		return mm
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return mm
	}

	for k, v := range m {
		mm[k] = v[0]
	}
	return mm
}

func GetFollowGet(link string) string {
	redirectClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 3 * time.Second,
	}
	request, _ := http.NewRequest("GET", link, nil)
	do, err := redirectClient.Do(request)
	if err != nil {
		return ""
	}
	defer do.Body.Close()
	body, err := io.ReadAll(do.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func AddParam(link string, param string) string {
	if strings.Contains(link, "?") {
		return link + "&" + param
	}
	return link + "?" + param
}
