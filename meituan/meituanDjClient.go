package meituan

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type MeiuanDjClient struct {
}

func NewMeituanDjClient() *MeiuanDjClient {
	return &MeiuanDjClient{}
}

func (m *MeiuanDjClient) Execute(params DjRequest) (string, error) {
	marshal, _ := json.Marshal(params.GetApiParamsObject())
	return m.HttpPostJson(params.GetApiPath(), string(marshal))
}

func (m *MeiuanDjClient) HttpPostJson(url string, data string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
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
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var DefaultMeituanDjClient *MeiuanDjClient = NewMeituanDjClient()
