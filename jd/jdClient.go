package jd

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type JdClient struct {
}

func NewJdClient() *JdClient {
	return &JdClient{}
}

func (j *JdClient) Request(params JdBaseApiRequest, isNeedAuth bool) (string, error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	buy_param_json, _ := json.Marshal(params.GetParamsObject())
	fields := map[string]string{
		"timestamp":         timeStr,
		"v":                 ApiVersion,
		"sign_method":       "md5",
		"format":            "json",
		"method":            params.GetMethodName(),
		"360buy_param_json": string(buy_param_json),
		"app_key":           params.GetConfig().AppKey,
	}
	accessToken := ""
	if isNeedAuth {
		fields["access_token"] = params.GetConfig().AccessToken
		accessToken = params.GetConfig().AccessToken
	}
	sign := j.sign(string(buy_param_json), accessToken, params.GetConfig().AppKey, "json", params.GetMethodName(), "md5", timeStr, ApiVersion, params.GetConfig().AppSecret)
	fields["sign"] = sign

	return HttpPost(params.GetConfig().BaseUrl, fields)
}

func (j *JdClient) sign(param_json, access_token, app_key, format, method, sign_method, timestamp, v, app_secret string) string {
	signStr := app_secret + "360buy_param_json" + param_json
	if access_token != "" {
		signStr += "access_token" + access_token
	}
	signStr += "app_key" + app_key + "format" + format + "method" + method + "sign_method" + sign_method + "timestamp" + timestamp + "v" + v + app_secret
	s := Md5(signStr)
	return strings.ToUpper(s)
}

func Md5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
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

var DefaultJdApiClient *JdClient = NewJdClient()
