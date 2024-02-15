package qianzhu

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/qianzhu/qianzhuResponse"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	KFC_URL          = "/kfc/?platformId=%s&token=%s"
	MOVIE_URL        = "/cinema/?platformId=%s&token=%s"
	ACCESS_TOKEN_URL = "/api/v2/platform/getToken?platformId={platformId}&platformUniqueId={platformUniqueId}&nickname={nickname}&mobile={mobile}&sign={sign}&timestamp={timestamp}"
	PAY_KFC_ORDER    = "/openApi/v1/orders/payKfcOrder"
	PAY_MOVIE_ORDER  = "/openApi/v1/orders/payMovieOrder"
	GET_KFC_ORDER    = "/openApi/v1/kfcOrders/getByOrderNo"
	GET_MOVIE_ORDER  = "/openApi/v1/orders/getByOrderNo"
)

type QianzhuClient struct {
	PlatformId string
	Secret     string
	Url        string
	H5Url      string
	MiniUrl    string
}

func NewQianzhuClient(platformId, secret, url, h5url, miniUrl string) *QianzhuClient {
	return &QianzhuClient{
		PlatformId: platformId,
		Secret:     secret,
		Url:        url,
		H5Url:      h5url,
		MiniUrl:    miniUrl,
	}
}

func (c *QianzhuClient) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = ""
	for k, v := range keys {
		ss += v + "=" + params[v]
		if k < len(keys)-1 {
			ss += "&"
		}
	}
	ss += c.Secret
	h := md5.New()
	h.Write([]byte(ss))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

func (c *QianzhuClient) getAccessToken(platformUniqueId, phone string) (string, error) {
	params := map[string]string{
		"platformUniqueId": platformUniqueId,
		"nickname":         "",
		"mobile":           phone,
		"timestamp":        cast.ToString(time.Now().Unix()),
		"platformId":       c.PlatformId}
	sign := c.generateSign(params)
	params["sign"] = sign

	urlStr := c.Url + ACCESS_TOKEN_URL
	for key, value := range params {
		v := url.QueryEscape(value)
		urlStr = strings.ReplaceAll(urlStr, "{"+key+"}", v)
	}

	resStr, httpErr := utils.GetResponse(urlStr, nil, 3*time.Second)

	var data qianzhuResponse.AccessTokenResponse
	if httpErr != nil {
		return "", httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return "", jsonErr
	}
	if data.Code == 10000 && data.Data.AccessToken != "" {
		return data.Data.AccessToken, nil
	}
	return "", errors.New("请稍后再试")
}

func (c *QianzhuClient) GetKfcUrl(platformUniqueId, phone string) (string, error) {
	accessToken, err := c.getAccessToken(platformUniqueId, phone)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf(c.H5Url+KFC_URL, c.PlatformId, accessToken)

	return url, nil
}

func (c *QianzhuClient) GetMovieUrl(platformUniqueId, phone string) (string, error) {
	accessToken, err := c.getAccessToken(platformUniqueId, phone)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf(c.H5Url+MOVIE_URL, c.PlatformId, accessToken)

	return url, nil
}

func (c *QianzhuClient) PayOrder(orderNo, orderUrl string) error {
	var params map[string]string
	if orderUrl == PAY_MOVIE_ORDER {
		params = map[string]string{
			"orderNo":    orderNo,
			"timestamp":  cast.ToString(time.Now().Unix()),
			"platformId": c.PlatformId}
	} else {
		params = map[string]string{
			"notifyUrl":  "",
			"orderNo":    orderNo,
			"timestamp":  cast.ToString(time.Now().Unix()),
			"platformId": c.PlatformId}
	}

	sign := c.generateSign(params)
	params["sign"] = sign

	urlStr := c.Url + orderUrl
	jsonStr, _ := json.Marshal(params)
	resStr, httpErr := utils.HttpPostJson(urlStr, string(jsonStr), time.Second*3, nil)

	var data qianzhuResponse.PayResponse
	if httpErr != nil {
		return httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return jsonErr
	}
	if data.Code == 10000 && data.Data.Balance != 0 {
		return nil
	}
	return nil
}

func (c *QianzhuClient) GetOrder(orderNo, orderUrl string) (qianzhuResponse.OrderDetailResponse, error) {
	params := map[string]string{
		"orderNo":    orderNo,
		"timestamp":  cast.ToString(time.Now().Unix()),
		"platformId": c.PlatformId}
	sign := c.generateSign(params)
	params["sign"] = sign

	urlStr := c.Url + orderUrl
	jsonStr, _ := json.Marshal(params)
	resStr, httpErr := utils.HttpPostJson(urlStr, string(jsonStr), time.Second*3, nil)

	var data qianzhuResponse.OrderDetailResponse
	if httpErr != nil {
		return data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(resStr), &data)
	if jsonErr != nil {
		return data, jsonErr
	}
	return data, nil
}
