package douyin_life

import (
	"bytes"
	"doudian.com/life/request"
	"doudian.com/life/response"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type Client struct {
	Domain       string
	ClientKey    string
	ClientSecret string
	Cache        Cache
}

func NewClient(clientKey, clientSecret string, cache Cache) *Client {
	return &Client{
		Domain:       "https://open.douyin.com",
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		Cache:        cache,
	}
}

type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}

type AccessTokenResp struct {
	Data    AccessTokenInfo `json:"data"`
	Message string          `json:"message"`
}
type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	Captcha     string `json:"captcha"`
	DescURL     string `json:"desc_url"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
	ExpiresIn   int    `json:"expires_in"`
	LogID       string `json:"log_id"`
}

const DouyinLifeAccessToken = "key:douyin:life:accessToken:"

func (c *Client) GetAccessToken() string {
	cacheKey := DouyinLifeAccessToken + c.ClientKey
	if c.Cache.IsExist(cacheKey) {
		accessToken := c.Cache.Get(cacheKey)
		if accessToken != "" && accessToken != nil {
			return fmt.Sprintf("%v", accessToken)
		}
	}

	accessToken := c.RefreshAccessToken()
	return accessToken
}

func (c *Client) RefreshAccessToken() string {
	cacheKey := DouyinLifeAccessToken + c.ClientKey
	m := make(map[string]string)
	m["grant_type"] = "client_credential"
	m["client_key"] = c.ClientKey
	m["client_secret"] = c.ClientSecret

	uri := c.Domain + "/oauth/client_token/"
	d, _ := json.Marshal(m)
	resp, err := c.request(uri, d, "")
	if err != nil {
		return ""
	}
	var accessTokenResp AccessTokenResp
	err = json.Unmarshal(resp, &accessTokenResp)
	if err != nil {
		return ""
	}
	accessTokenStr := accessTokenResp.Data.AccessToken
	c.Cache.Set(cacheKey, accessTokenStr, 7200*time.Second)
	return accessTokenStr
}

func (c *Client) CommandParseAndShare(req *request.CommandParseAndShareRequest) (*response.CommandParseAndShareResponse, error) {
	var res response.CommandParseAndShareResponse
	i, err := c.doRequestAndResponse(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(i, &res)
	if err != nil {
		return nil, err
	}
	if res.ErrNo != 0 {
		return nil, errors.New(res.ErrMsg)
	}
	return &res, nil
}

func (c *Client) PidCreate(req *request.PidCreateRequest) (*response.PidCreateResponse, error) {
	var res response.PidCreateResponse
	i, err := c.doRequestAndResponse(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(i, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ProductSearch(req *request.SearchProductRequest) (*response.SearchProductResponse, error) {
	var res response.SearchProductResponse
	i, err := c.doRequestAndResponse(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(i, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ProductDetail(req *request.ProductDetailRequest) (*response.ProductDetailResponse, error) {
	var res response.ProductDetailResponse
	i, err := c.doRequestAndResponse(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(i, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) doRequestAndResponse(req ApiRequest) ([]byte, error) {
	accessToken := c.GetAccessToken()
	uri := c.Domain + req.GetUrlPath()
	data, err := json.Marshal(req.GetParams())
	d, err := c.request(uri, data, accessToken)
	if err != nil {
		return nil, err
	}
	errNo := gjson.Get(string(d), "err_no").Int()
	if errNo == 28001008 {
		accessToken = c.RefreshAccessToken()
		d, err = c.request(uri, data, accessToken)
		if err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (c *Client) request(uri string, data []byte, accessToken string) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil && req == nil {
		return nil, err
	}

	if req == nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
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
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	if accessToken != "" {
		req.Header.Set("access-token", accessToken)
	}
	//发送请求
	resp, err := client.Do(req)
	if resp == nil {
		return nil, err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
