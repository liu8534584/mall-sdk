package pangolin

import (
	"fmt"
	"github.com/polaris1119/goutils"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func getSign(reqBody ReqBody, securityKey string) string {

	signStr := fmt.Sprintf("app_id=%v&data=%v&req_id=%v&timestamp=%d%v", reqBody.AppId, reqBody.Data, reqBody.ReqId, reqBody.Timestamp, securityKey)

	md5 := goutils.Md5(signStr)

	return md5
}

func PangolinPost(url string, data string) (string, string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", "", err
	}

	if req == nil {
		return "", "", err
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
		return "", "", err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return "", "", err
	}
	logId := resp.Header.Get("X-Tt-Logid")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	return string(body), logId, nil
}
