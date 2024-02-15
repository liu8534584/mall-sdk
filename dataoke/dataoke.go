package dataoke

import (
	"encoding/json"
	"errors"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/polaris1119/goutils"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

const (
	BaseUrl = "https://openapi.dataoke.com"
	Version = "v1.3.0"
)

type client struct {
	AppKey    string
	AppSecret string
	Timeout   time.Duration
}

func NewDataokeClient(appKey, appSecret string, timeout time.Duration) *client {
	return &client{
		AppKey:    appKey,
		AppSecret: appSecret,
		Timeout:   timeout,
	}
}

func (c *client) Timout(timeout time.Duration) {
	c.Timeout = timeout
}

// GetRankList 获取榜单列表
func (c *client) GetRankList(rankType, cid, pageNo, pageSize int) (rankList []RankList, err error) {
	data := map[string]string{
		"rankType": cast.ToString(rankType),
		"cid":      cast.ToString(cid),
		"pageSize": cast.ToString(pageSize),
		"pageId":   cast.ToString(pageNo),
	}
	method := "/api/goods/get-ranking-list"
	rsp, err := c.doRequest(method, data)

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(rsp.Data)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonData, &rankList)
	if err != nil {
		return
	}
	return
}

//发送请求
func (c *client) doRequest(method string, data map[string]string) (Response, error) {
	urlPath := BaseUrl + method
	data["version"] = Version
	data["appKey"] = c.AppKey
	data["sign"] = c.createSign(data)
	urlPath = utils.SetQuery(urlPath, data)
	resp, err := utils.GetResponse(urlPath, nil, c.Timeout)
	var res Response
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	if res.Code != 0 {
		return res, errors.New(res.Msg)
	}
	return res, nil
}

func (c *client) createSign(data map[string]string) string {
	var keys []string
	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var md5Str string
	md5Str = ""
	for _, kk := range keys {
		md5Str += "&" + string(kk) + "=" + cast.ToString(data[kk])
	}
	md5Str = strings.Trim(md5Str, "&") + "&key=" + c.AppSecret
	return strings.ToUpper(goutils.Md5(md5Str))
}
