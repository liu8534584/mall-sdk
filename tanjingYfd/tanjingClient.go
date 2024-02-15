package tanjingYfd

import (
	"encoding/json"
	"github.com/spf13/cast"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type TanjingClient struct {
}

func NewTanjingClient() *TanjingClient {
	return &TanjingClient{}
}

func (t *TanjingClient) Execute(request TanjingRequest, accessToken string) (string, error) {
	requestUrl := request.GetConfig().BaseUrl + request.GetApiPath()
	marshal, _ := json.Marshal(request.GetApiParamsObject())
	return t.HttpPostJson(requestUrl, string(marshal), accessToken, request.GetConfig().Timeout)
}

func (t *TanjingClient) HttpPostJson(url string, data string, token string, timeout int64) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
	}

	client := &http.Client{
		Timeout: time.Duration(timeout * cast.ToInt64(time.Second)),
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3)
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
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("token", token)
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

var DefaultTanjingClient *TanjingClient = NewTanjingClient()
