package czb

import (
	"context"
	"encoding/json"
	"github.com/liu8534584/mall-sdk/czb/request"
	"github.com/liu8534584/mall-sdk/czb/response"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/polaris1119/goutils"
	"github.com/spf13/cast"
	"sort"
	"time"
)

type Client struct {
	Timeout     time.Duration
	OrderSource string //渠道编码
	AppKey      string
	AppSecret   string
	BaseUrl     string
}

func NewClient(appKey, appSecret, orderSource string, timeout time.Duration) *Client {
	return &Client{
		Timeout:     timeout,
		OrderSource: orderSource,
		AppKey:      appKey,
		AppSecret:   appSecret,
		BaseUrl:     "https://mcs.czb365.com",
	}
}

func (c *Client) GetSecretCode(ctx context.Context, req request.GetSecretCodeRequest) (response.GetSecretCodeResponse, error) {
	var resp response.GetSecretCodeResponse
	err := c.DoRequest(ctx, req.GetParams(), req.GetUrlPath(), &resp)
	if err != nil {
		return response.GetSecretCodeResponse{}, err
	}
	return resp, nil
}

func (c *Client) DoRequest(ctx context.Context, params map[string]interface{}, urlPath string, resp any) error {

	params["app_key"] = c.AppKey
	params["timestamp"] = time.Now().UnixMilli()
	params["sign"] = c.getSign(params)

	jsonStr, _ := json.Marshal(params)
	urlPath = c.BaseUrl + urlPath
	respStr, err := utils.HttpPostJson(urlPath, string(jsonStr), c.Timeout, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		return err
	}
	return nil

}
func (c *Client) getSign(params map[string]interface{}) string {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	keyStr := c.AppSecret
	for _, v := range keys {
		keyStr += v + cast.ToString(params[v])
	}
	keyStr += c.AppSecret
	return goutils.Md5(keyStr)
}
