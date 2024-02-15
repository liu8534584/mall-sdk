package kuaishou

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/kuaishou/kuaishouapi"
	"github.com/liu8534584/mall-sdk/utils"
	"strconv"
	"strings"
	"time"
)

type KuaishouClient struct {
	appKey     string
	signSecret string
	gatewayUrl string
	CpsPid     string
}

func NewKuaishouClient(appKey, signSecret, CpsPid string) *KuaishouClient {
	return &KuaishouClient{
		appKey:     appKey,
		signSecret: signSecret,
		gatewayUrl: "https://openapi.kwaixiaodian.com/",
		CpsPid:     CpsPid,
	}
}

func (c *KuaishouClient) execute(method, accessToken string, params interface{}, version int, isPost bool) (string, error) {

	var sysParams = make(map[string]string)
	sysParams["appkey"] = c.appKey
	timestamp := time.Now().Unix() * 1000
	sysParams["timestamp"] = strconv.FormatInt(timestamp, 10)
	sysParams["access_token"] = accessToken
	sysParams["version"] = strconv.Itoa(version)
	paramJson, sign := c.sign(accessToken, method, timestamp, version, params)
	sysParams["sign"] = sign
	sysParams["method"] = method
	sysParams["signMethod"] = "MD5"

	apiUrl := c.gatewayUrl + strings.ReplaceAll(method, ".", "/")
	var err error
	var body string
	if isPost {
		sysParams["param"] = paramJson
		body, err = utils.HttpPost(apiUrl, sysParams)
	} else {
		sysParams["param"] = paramJson
		return utils.HttpGet(utils.SetQuery(apiUrl, sysParams), 3*time.Second)
	}
	return body, err
}

func (c *KuaishouClient) ItemList(ctx context.Context, request kuaishouapi.ItemListRequest, accessToken string) (kuaishouapi.ItemListResponseData, error) {

	var res kuaishouapi.ItemListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodItemList, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) ItemDetail(ctx context.Context, request kuaishouapi.ItemDetailRequest, accessToken string) ([]kuaishouapi.KuaishouItem, error) {

	var res kuaishouapi.ItemDetailResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodItemDetail, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) LinkCreate(ctx context.Context, request kuaishouapi.LinkCreateRequest, accessToken string) (kuaishouapi.LinkCreateResponseData, error) {

	var res kuaishouapi.LinkCreateResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodLinkCreate, accessToken, request, 1, true)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) ChannelList(ctx context.Context, accessToken string) ([]kuaishouapi.ChannelListResponseChannel, error) {

	var res kuaishouapi.ChannelListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodChannelList, accessToken, map[string]string{}, 1, false)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, nil
}

func (c *KuaishouClient) PromtionThemeItemList(ctx context.Context, request kuaishouapi.PromtionThemeItemListRequest, accessToken string) ([]kuaishouapi.KuaishouItem, string, error) {

	var res kuaishouapi.PromtionThemeItemListResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodPromtionThemeItemList, accessToken, request, 1, false)
	if httpErr != nil {
		return res.Data, res.Pcursor, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, res.Pcursor, jsonErr
	}

	if res.Result != 1 {
		return res.Data, res.Pcursor, NewApiError(res.Result, fmt.Sprintf("message:%s,code:%s,subMessage:%s,subCode:%s,err:%v", res.Msg, res.Code, res.SubMsg, res.SubCode, res))
	}

	return res.Data, res.Pcursor, nil
}

func (c *KuaishouClient) LinkParse(ctx context.Context, request kuaishouapi.LinkParseRequest, accessToken string) (kuaishouapi.LinkParseResponseData, error) {

	var res kuaishouapi.LinkParseResponse
	httpExecute, httpErr := c.execute(kuaishouapi.MethodLinkParse, accessToken, request, 1, true)
	if httpErr != nil {
		return res.Data, httpErr
	}

	jsonErr := json.Unmarshal([]byte(httpExecute), &res)
	if jsonErr != nil {
		return res.Data, jsonErr
	}

	if res.Result != 1 {
		return res.Data, errors.New("快手内部错误,请稍后再试")
	}

	return res.Data, nil
}
