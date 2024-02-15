package vip

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/spf13/cast"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

type VipClient struct {
	Timeout time.Duration
}

func NewVipClient(timeout time.Duration) *VipClient {
	return &VipClient{Timeout: timeout}
}

func (j *VipClient) Request(params VipBaseApiRequest, isNeedAuth bool) (string, error) {
	timeStr := cast.ToString(time.Now().Unix())
	service := params.GetServiceName()
	method := params.GetMethodName()
	format := "json"
	busiParams, _ := json.Marshal(params.GetParamsObject())

	accessToken := ""
	if isNeedAuth {
		accessToken = params.GetConfig().AccessToken
	}
	sign := j.createRequestSign(accessToken, params.GetConfig().AppKey, format, "", method, service, timeStr, ApiVersion, string(busiParams), params.GetConfig().AppSecret)
	fields := map[string]string{
		"appKey":    params.GetConfig().AppKey,
		"format":    format,
		"method":    method,
		"service":   service,
		"sign":      sign,
		"timestamp": timeStr,
		"version":   ApiVersion,
	}
	if accessToken != "" {
		fields["accessToken"] = accessToken
	}
	uri := params.GetConfig().BaseUrl + "?" + getQueryString(fields)
	return HttpPost(uri, string(busiParams), j.Timeout)
}

func (j *VipClient) createRequestSign(access_token, appKey, format, language, method, service, timestamp, version, busiParams, appSecret string) string {
	signStr := ""
	if access_token != "" {
		signStr += "accessToken" + access_token
	}
	signStr += "appKey" + appKey + "format" + format
	if language != "" {
		signStr += "language" + language
	}
	signStr += "method" + method + "service" + service + "timestamp" + timestamp + "version" + version + busiParams

	return Hmac(signStr, appSecret)
}

func getQueryString(maps map[string]string) string {
	queryString := ""
	i := 0
	for k, v := range maps {
		if i == 0 {
			queryString += k + "=" + v
		} else {
			queryString += "&" + k + "=" + v
		}
		i++
	}

	return queryString
}

func Hmac(data, key string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))

	return strings.ToUpper(hex.EncodeToString(hmac.Sum([]byte(""))))
}

func HttpPost(urls string, data string, timeout time.Duration) (string, error) {
	var c *http.Client = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, timeout)
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

	resp, err := c.Post(urls, "application/json,charset=utf-8",
		strings.NewReader(data))
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

var DefaultVipApiClient *VipClient = NewVipClient(3 * time.Second)
