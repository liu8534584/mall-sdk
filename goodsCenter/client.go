package goodsCenter

import (
	"context"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/liu8534584/mall-sdk/goodsCenter/response/base"
	"time"
)

type APIError struct {
	Code    base.Code `json:"code"`
	Message string    `json:"msg"`
}

func (e *APIError) Error() string {
	msg := fmt.Sprintf("API error: %s", e.Message)
	return fmt.Sprintf("%s ", msg)
}

type Client struct {
	*req.Client
}

func NewGoodsCenterClient(ctx context.Context, envName string, timeout time.Duration, appName string) *Client {
	baseUrl := "http://172.16.28.209"
	if envName == "prerelease" || envName == "develop" {
		baseUrl = "http://127.0.0.1:2600"
	}
	return &Client{
		Client: req.C().
			SetBaseURL(baseUrl).
			SetCommonError(&APIError{}).
			AddCommonQueryParam("app", appName).
			SetTimeout(timeout).
			SetCommonRetryCount(2).
			EnableDumpEachRequest().
			OnAfterResponse(func(client *req.Client, resp *req.Response) error {
				if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
					return errors.New("请稍后再试")
				}
				if apiErr, ok := resp.Error().(*APIError); ok {
					// Server returns an error message, convert it to human-readable go error.
					resp.Err = apiErr
					return nil
				}
				// Corner case: neither an error state response nor a success state response,
				// dump content to help troubleshoot.
				if !resp.IsSuccess() {
					return fmt.Errorf("bad response, raw dump:\n%s", resp.Dump())
				}
				return nil
			}),
	}
}

func (c *Client) SetDebug(enable bool) *Client {
	if enable {
		c.EnableDebugLog()
		c.EnableDumpAll()
	} else {
		c.DisableDebugLog()
		c.DisableDumpAll()
	}
	return c
}
