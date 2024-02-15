package Kuaizhan

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"sort"
	"time"
)

type KuaizhanClient struct {
	Key        string
	Secret     string
	gatewayUrl string
	Timeout    time.Duration
}

type MtActTransParam struct {
	SId         string `json:"Sid"`
	ActId       string `json:"actId"`
	LinkType    int    `json:"linkType"`
	IsShortLink int    `json:"isShortLink"`
}

func NewKuaizhanClient(key, secret string, timeout time.Duration) *KuaizhanClient {
	return &KuaizhanClient{
		Key:        key,
		Secret:     secret,
		gatewayUrl: "https://cloud.kuaizhan.com/api/v1/km/genShortLink",
		Timeout:    timeout,
	}
}

func (c *KuaizhanClient) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = c.Secret
	for _, v := range keys {
		ss += v + params[v]
	}
	ss += c.Secret
	h := md5.New()
	h.Write([]byte(ss))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *KuaizhanClient) execute(params map[string]interface{}, isPost bool) (string, error) {
	var sysParams = make(map[string]string)
	sysParams["appKey"] = c.Key
	sysParams["format"] = "json"
	for k, v := range params {
		sysParams[k] = cast.ToString(v)
	}
	sysParams["sign"] = c.generateSign(sysParams)

	if isPost {
		return utils.HttpPost(c.gatewayUrl, sysParams)
	}

	response, err := utils.GetResponse(utils.SetQuery(c.gatewayUrl, sysParams), nil, c.Timeout)
	if err != nil {
		return "", err
	}
	return string(response), nil

}

func (c *KuaizhanClient) GenShortUrlResp(url string) (GenShortUrlResp, error) {

	params := map[string]interface{}{"link": url}
	resStr, httpErr := c.execute(params, false)

	var data GenShortUrlResp
	if httpErr != nil {
		return data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return data, jsonErr
	}

	if data.Code != 0 {
		return data, errors.New("快站内部错误,请稍后再试")
	}

	return data, nil
}
