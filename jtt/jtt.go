package jtt

import (
	"encoding/json"
	"errors"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"time"
)

const (
	BaseUrl = "http://japi.jingtuitui.com"
	Version = "v2"
)

type client struct {
	AppId     string
	AppKey    string
	AppSecret string
	Timeout   time.Duration
}

func NewJttClient(appId, appKey string, timeout time.Duration) *client {
	return &client{
		AppKey:  appKey,
		AppId:   appId,
		Timeout: timeout,
	}
}

func (c *client) Timout(timeout time.Duration) {
	c.Timeout = timeout
}

// GetRankList 获取榜单列表
func (c *client) GetRankList(cid string, pageNo, pageSize int) (rankList []RankList, err error) {
	data := map[string]string{
		"eliteId":   cid,
		"pageSize":  cast.ToString(pageSize),
		"pageIndex": cast.ToString(pageNo),
	}
	method := "/api/today_top"
	rsp, err := c.doRequest(method, data)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(rsp.Result)
	if err != nil {
		return
	}
	itemList := gjson.Get(string(jsonData), "data").String()
	err = json.Unmarshal([]byte(itemList), &rankList)
	if err != nil {
		return
	}
	return
}

// 发送请求
func (c *client) doRequest(method string, data map[string]string) (Response, error) {
	urlPath := BaseUrl + method
	data["appid"] = c.AppId
	data["appkey"] = c.AppKey
	data["v"] = Version
	urlPath = utils.SetQuery(urlPath, data)
	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
	resp, err := utils.GetResponse2(urlPath, headers, c.Timeout, 0)
	var res Response
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return res, err
	}
	if res.Return != 0 {
		return res, errors.New(cast.ToString(res.Result))
	}
	return res, nil
}
