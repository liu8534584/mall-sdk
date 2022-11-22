package meituan

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanResponse"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

type meituanClient struct {
	Key        string
	Secret     string
	gatewayUrl string
	Timeout    time.Duration
}

func NewMeituanClient(key, secret string, timeout time.Duration) *meituanClient {
	return &meituanClient{
		Key:        key,
		Secret:     secret,
		gatewayUrl: "https://openapi.meituan.com/api/",
		Timeout:    timeout,
	}
}

func (c *meituanClient) generateSign(params map[string]string) string {
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

func (c *meituanClient) execute(method string, params map[string]interface{}, isPost bool) (string, error) {
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

func (c *meituanClient) GenerateLink(linkType int, actId string, userId uint64) (meituanResponse.MeituanGenerateLinkResponse, error) {

	params := map[string]interface{}{
		"actId":    actId,
		"sid":      fmt.Sprintf("lmm%d", userId),
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

func (c *meituanClient) MiniCode(linkType int, actId string, userId uint64) (meituanResponse.MeituanGenerateLinkResponse, error) {
	params := map[string]interface{}{
		"actId":    actId,
		"sid":      fmt.Sprintf("lmm%d", userId),
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
