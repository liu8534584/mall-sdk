package suning

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/liu8534584/mall-sdk/suning/request"
	"github.com/liu8534584/mall-sdk/suning/response"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/polaris1119/goutils"
	"github.com/spf13/cast"
	"time"
)

type SuningClient struct {
	AppKey     string
	AppSecret  string
	ApiVersion string
	ServerUrl  string
	Timeout    time.Duration
}

type snRequest struct {
	SnBody map[string]interface{} `json:"sn_body"`
}

type Request struct {
	SnRequest snRequest `json:"sn_request"`
}

type sysParams struct {
	SecretKey  string `json:"secret_key"`
	Method     string `json:"method"`
	Date       string `json:"date"`
	AppKey     string `json:"app_key"`
	ApiVersion string `json:"api_version"`
	PostField  string `json:"post_field"`
}

func NewSuningClient(appKey string, appSecret string) *SuningClient {
	return &SuningClient{
		AppKey:     appKey,
		AppSecret:  appSecret,
		ApiVersion: "v1.2",
		ServerUrl:  "https://open.suning.com/api/http/sopRequest",
		Timeout:    2 * time.Second,
	}
}

func (c *SuningClient) GetOrderList(orderList request.OrderListReq) (response.GetOrderListResp, error) {
	params := map[string]interface{}{
		"endTime":         orderList.EndTime,
		"startTime":       orderList.StartTime,
		"pageNo":          cast.ToString(orderList.PageNo),
		"pageSize":        cast.ToString(orderList.PageSize),
		"orderLineStatus": cast.ToString(orderList.OrderLineStatus),
	}

	method := "suning.netalliance.order.query"
	queryName := "queryOrder"
	jsonData := c.doRequest(method, queryName, params)
	var res response.GetOrderListResp
	err := json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		return res, err
	}
	if len(res.SnResponseContent.SnBody.QueryOrder) == 0 {
		return res, errors.New(fmt.Sprintf("查询失败,err:%v", res))
	}
	return res, nil
}

func (c *SuningClient) GetOrderByOrderNo(orderNo string) (response.OrderInfoResp, error) {
	params := map[string]interface{}{
		"orderCode": orderNo,
	}

	method := "suning.netalliance.order.get"
	queryName := "getOrder"
	jsonData := c.doRequest(method, queryName, params)
	var res response.OrderInfoResp
	err := json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetItemInfo 获取商品信息
func (c *SuningClient) GetItemInfo(itemId, shopId string) (response.QueryCommodityDetailResponse, error) {
	shopId = utils.StrPad(shopId, 10, "0", 0)

	params := map[string]interface{}{
		"commodityStr": itemId + "-" + shopId,
		"picWidth":     400,
		"picHeight":    400,
		"couponMark":   "1",
	}

	method := "suning.netalliance.commoditydetail.query"
	queryName := "queryCommoditydetail"
	jsonData := c.doRequest(method, queryName, params)
	var res response.QueryCommodityDetailResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		return res, err
	}
	if res.SnResponseContent.SnBody.QueryCommoditydetail != nil {
		return res, nil
	}
	return res, errors.New(fmt.Sprintf("查询失败,err:%v", res))
}

func (c *SuningClient) GetRecommendItemsNew(page, pageSize, eliteId int64) (response.QuerySelectrecommendcommodityResponse, error) {
	params := map[string]interface{}{
		"pageIndex": page,
		"size":      pageSize,
		"eliteId":   eliteId,
		"picWidth":  400,
		"picHeight": 400,
	}
	method := "suning.netalliance.selectrecommendcommodity.query"
	queryName := "querySelectrecommendcommodity"
	jsonData := c.doRequest(method, queryName, params)
	var res response.QuerySelectrecommendcommodityResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

func (c *SuningClient) SearchItems(keyword string, sortType, page, pageSize int64) (response.QuerySearchcommodityResponse, error) {
	params := map[string]interface{}{
		"pageIndex": page,
		"size":      pageSize,
		"keyword":   keyword,
		"sortType":  sortType,
		"picWidth":  400,
		"picHeight": 400,
	}
	method := "suning.netalliance.searchcommodity.query"
	queryName := "querySearchcommodity"
	var res response.QuerySearchcommodityResponse
	jsonData := c.doRequest(method, queryName, params)
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

// QueryCommodityDetail 批量查询联盟商品信息
func (c *SuningClient) QueryCommodityDetail(commodityStr string) (response.QueryCommodityDetailResponse, error) {
	params := map[string]interface{}{
		"commodityStr": commodityStr,
	}
	method := "suning.netalliance.commoditydetail.query"
	queryName := "queryCommoditydetail"
	jsonData := c.doRequest(method, queryName, params)
	var res response.QueryCommodityDetailResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		return res, err
	}
	if res.SnResponseContent.SnBody.QueryCommoditydetail != nil {
		return res, nil
	}
	return res, errors.New(fmt.Sprintf("查询失败,err:%v", res))
}

func (c *SuningClient) GetEnterPriseyfgproUrl(actCode, actType, subUser string) (response.GetEnterPriseyfgproUrlResponse, error) {
	params := map[string]interface{}{
		"actCode": actCode,
		"proType": 0,
		"subUser": subUser,
		"actType": actType,
	}
	method := "suning.netalliance.enterpriseyfgprourl.get"
	queryName := "getEnterpriseyfgprourl"
	jsonData := c.doRequest(method, queryName, params)
	var res response.GetEnterPriseyfgproUrlResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

// GetItemUrl 转链
func (c *SuningClient) GetItemUrl(itemUrl, couponUrl, subUser, postionId string) (response.GetExtensionlinkResponse, error) {
	params := map[string]interface{}{
		"productUrl":  itemUrl,
		"promotionId": postionId,
	}

	if couponUrl != "" {
		params["quanUrl"] = couponUrl
	}

	if subUser != "" {
		params["subUser"] = subUser
	}
	method := "suning.netalliance.extensionlink.get"
	queryName := "getExtensionlink"
	jsonData := c.doRequest(method, queryName, params)
	var res response.GetExtensionlinkResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

// GetAppletextensionlink 获取小程序链接
func (c *SuningClient) GetAppletextensionlink(itemUrl, couponUrl, subUser, postionId string) (response.GetAppletextensionlinkResponse, error) {
	params := map[string]interface{}{
		"productUrl": itemUrl,
		"subUser":    subUser,
		"quanUrl":    couponUrl,
	}
	if postionId != "" {
		params["promotionId"] = postionId
	}

	method := "suning.netalliance.appletextensionlink.get"
	queryName := "getAppletextensionlink"
	jsonData := c.doRequest(method, queryName, params)
	var res response.GetAppletextensionlinkResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

func (c *SuningClient) Bacthcustomlink(extend, subUser, pid, promotionId string) (response.QueryBacthcustomlinkResponse, error) {
	params := map[string]interface{}{
		"extend":  extend,
		"subUser": subUser,
		"pid":     pid,
	}
	if promotionId != "" {
		params["promotionId"] = promotionId
	}

	method := "suning.netalliance.bacthcustomlink.query"
	queryName := "queryBacthcustomlink"
	jsonData := c.doRequest(method, queryName, params)
	var res response.QueryBacthcustomlinkResponse
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

// DoRequest 发起请求
func (c *SuningClient) doRequest(method string, queryName string, params map[string]interface{}) string {
	t := time.Now().Format("2006-01-02 15:04:05")

	snBody := map[string]interface{}{
		queryName: params,
	}
	field := Request{SnRequest: snRequest{SnBody: snBody}}
	fieldByte, _ := json.Marshal(field)
	postField := sysParams{
		SecretKey:  c.AppSecret,
		Method:     method,
		Date:       t,
		AppKey:     c.AppKey,
		ApiVersion: c.ApiVersion,
		PostField:  base64.StdEncoding.EncodeToString(fieldByte),
	}

	headers := c.generateSignHeader(postField)
	uri := c.ServerUrl + "/" + method
	res, err := grequests.Post(uri, &grequests.RequestOptions{
		JSON:    field,
		Headers: headers,
	})
	if err != nil {
		return ""
	}

	return res.String()

}

// generateSignHeader 生成sign header
func (c *SuningClient) generateSignHeader(fields sysParams) map[string]string {
	signStr := fields.SecretKey + fields.Method + fields.Date + fields.AppKey + fields.ApiVersion + fields.PostField
	signStr = goutils.Md5(signStr)

	headers := map[string]string{
		"Content-Type":   "application/json; charset=utf-8",
		"AppMethod":      fields.Method,
		"AppRequestTime": fields.Date,
		"Format":         "json",
		"signInfo":       signStr,
		"AppKey":         fields.AppKey,
		"VersionNo":      fields.ApiVersion,
		"User-Agent":     "suning-sdk-php",
		"Sdk-Version":    "suning-sdk-php-beta0.1",
	}
	return headers
}
