package meituan

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/liu8534584/mall-sdk/meituan/meituanResponse"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

type MeituanClient struct {
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

func NewMeituanClient(key, secret string, timeout time.Duration) *MeituanClient {
	return &MeituanClient{
		Key:        key,
		Secret:     secret,
		gatewayUrl: "https://openapi.meituan.com/api/",
		Timeout:    timeout,
	}
}

func (c *MeituanClient) generateSign(params map[string]string) string {
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
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func (c *MeituanClient) execute(method string, params map[string]interface{}, isPost bool) (string, error) {
	var sysParams = make(map[string]string)
	sysParams["appkey"] = c.Key
	sysParams["ts"] = cast.ToString(time.Now().Unix())
	for k, v := range params {
		sysParams[k] = cast.ToString(v)
	}
	sysParams["sign"] = c.generateSign(sysParams)

	if isPost {
		return utils.HttpPost(c.gatewayUrl+method, sysParams)
	}

	response, err := utils.GetResponse(utils.SetQuery(c.gatewayUrl+method, sysParams), nil, c.Timeout)
	if err != nil {
		return "", err
	}
	return string(response), nil

}

func (c *MeituanClient) GenerateLinkV2(param MtActTransParam) (meituanResponse.MeituanGenerateLinkResponse, error) {

	params := map[string]interface{}{
		"actId":     param.ActId,
		"sid":       param.SId,
		"shortLink": param.IsShortLink,
		"linkType":  param.LinkType}
	resStr, httpErr := c.execute("generateLink", params, false)

	var data meituanResponse.MeituanGenerateLinkResponse
	if httpErr != nil {
		return data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return data, jsonErr
	}
 
	if data.Status != 0 {
		return data, errors.New("美团内部错误,请稍后再试")
	}

	return data, nil
}

func (c *MeituanClient) GenerateLink(linkType int, actId string, sid string) (meituanResponse.MeituanGenerateLinkResponse, error) {

	params := map[string]interface{}{
		"actId":    actId,
		"sid":      sid,
		"linkType": linkType}
	resStr, httpErr := c.execute("generateLink", params, false)

	var data meituanResponse.MeituanGenerateLinkResponse
	if httpErr != nil {
		return data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return data, jsonErr
	}

	if data.Status != 0 {
		return data, errors.New("美团内部错误,请稍后再试")
	}

	return data, nil
}

func (c *MeituanClient) MiniCode(linkType int, actId string, sid string) (meituanResponse.MeituanGenerateLinkResponse, error) {
	params := map[string]interface{}{
		"actId":    actId,
		"sid":      sid,
		"linkType": linkType,
	}
	resStr, httpErr := c.execute("miniCode", params, false)

	var data meituanResponse.MeituanGenerateLinkResponse
	if httpErr != nil {
		return data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return data, jsonErr
	}

	if data.Status != 0 {
		return data, errors.New("美团内部错误,请稍后再试")
	}

	return data, nil
}
