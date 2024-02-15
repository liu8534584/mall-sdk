package kaola

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"net/url"
	"sort"
	"strings"
	"time"
)

type KaolaClient struct {
	ServerUrl   string
	AccessToken string
	AppKey      string
	AppSecret   string
	Timeout     time.Duration
}

func NewKaolaClient(serverUrl, accessToken, appKey, appSecret string, timeout time.Duration) *KaolaClient {
	return &KaolaClient{
		ServerUrl:   serverUrl,
		AccessToken: accessToken,
		AppKey:      appKey,
		AppSecret:   appSecret,
		Timeout:     timeout,
	}
}

func (k *KaolaClient) DoPost(params KaolaBaseRequest, isPost bool) (string, error) {

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fields := map[string]string{
		"timestamp":  timeStr,
		"v":          apiVersion,
		"signMethod": "md5",
		"method":     params.GetApiMethodName(),
		"unionId":    k.AppKey,
		"charset":    "utf-8",
	}
	for k, v := range params.GetApiParams() {
		fields[k] = cast.ToString(v)
	}
	fields["sign"] = k.generateSign(fields)

	if isPost {
		var clusterInfo = url.Values{}
		for k, v := range fields {
			clusterInfo.Add(k, v)
		}
		encode := clusterInfo.Encode()
		return utils.HttpPostJson(k.ServerUrl, encode, 3*time.Second, map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Accept":       "application/json; charset=utf-8",
		})
	}

	response, err := utils.GetResponse(utils.SetQuery(k.ServerUrl, fields), nil, k.Timeout)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (k *KaolaClient) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = k.AppSecret
	for _, v := range keys {
		ss += v + params[v]
	}
	ss += k.AppSecret
	h := md5.New()
	h.Write([]byte(ss))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
