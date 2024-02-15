package tencentCloud

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/liu8534584/mall-sdk/tencentCloud/request"
	"github.com/liu8534584/mall-sdk/tencentCloud/response"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"io"
	"sort"
	"time"
)

type TencentCloudClient struct {
	AppIdV1     string
	AppSecretV1 string
	AppIdV2     string
	AppSecretV2 string
	BaseUrlV1   string
	BaseUrlV2   string
	Timeout     time.Duration
	ShareId     string
}

func NewTencentCloudClient(appIdv1, appSecretv1, appIdv2, appSecretv2, shareId string, timeout time.Duration) *TencentCloudClient {
	return &TencentCloudClient{
		AppIdV1:     appIdv1,
		AppSecretV1: appSecretv1,
		AppIdV2:     appIdv2,
		AppSecretV2: appSecretv2,
		ShareId:     shareId,
		BaseUrlV1:   "https://zhls.qq.com",
		BaseUrlV2:   "https://openapi.ym.qq.com/gateway",
		Timeout:     3 * time.Second,
	}
}

func (c *TencentCloudClient) doRequestV2(api string, data map[string]interface{}) (string, error) {
	sysParams := map[string]string{
		"sropAppKey":     c.AppIdV2,
		"sropApiVersion": "1.0.0",
		"sropVersion":    "1.0.0",
		"sropTimestamp":  cast.ToString(time.Now().Unix() * 1000),
		"sropApi":        api,
	}

	requestBody, _ := json.Marshal(data)
	sysParams["sropSign"] = c.createSign(sysParams, string(requestBody))

	uri := c.BaseUrlV2
	var res *grequests.Response
	var err error
	uri = utils.SetQuery(uri, sysParams)
	headers := map[string]string{
		"Content-type": "application/json;charset='utf-8'",
	}
	res, err = grequests.Post(uri, &grequests.RequestOptions{
		JSON:           requestBody,
		Headers:        headers,
		RequestTimeout: c.Timeout,
	})
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (c *TencentCloudClient) createSign(params map[string]string, requestBody string) string {
	params["sropRaw"] = requestBody
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	query := c.AppSecretV2

	for k, v := range keys {
		if k == 0 {
			query += v + "=" + cast.ToString(params[v])
		} else {
			query += "&" + v + "=" + cast.ToString(params[v])
		}
	}
	query += c.AppSecretV2

	h := md5.New()
	_, _ = io.WriteString(h, query)
	return hex.EncodeToString(h.Sum(nil))
}

func (c *TencentCloudClient) doRequestV1(path string, data map[string]string, method string) (string, error) {
	sysParams := map[string]string{
		"app_id":    c.AppIdV1,
		"nonce":     uuid.NewV4().String(),
		"sign":      "sha256",
		"timestamp": cast.ToString(time.Now().Unix()),
	}

	sysParams["signature"] = c.createSignV1(sysParams)

	uri := c.BaseUrlV1 + path
	var res *grequests.Response
	var err error
	if method == "post" {
		uri = utils.SetQuery(uri, sysParams)
		headers := map[string]string{
			"Content-type": "application/json;charset='utf-8'",
		}
		res, err = grequests.Post(uri, &grequests.RequestOptions{
			JSON:           data,
			Headers:        headers,
			RequestTimeout: c.Timeout,
		})
	} else {
		data = utils.NewArrayUtils().Merge(sysParams, data)
		uri = utils.SetQuery(uri, data)
		res, err = grequests.Get(uri, &grequests.RequestOptions{RequestTimeout: c.Timeout})
	}
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (c *TencentCloudClient) createSignV1(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = ""
	for _, v := range keys {
		if ss == "" {
			ss += v + "=" + params[v]
		} else {
			ss += "&" + v + "=" + params[v]
		}
	}
	h := hmac.New(sha256.New, []byte(c.AppSecretV1))
	_, _ = h.Write([]byte(ss))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *TencentCloudClient) GetSkuDetail(spuIdList, saasId string) (response.CpsSkuDetail, error) {
	params := map[string]interface{}{
		"parameter": map[string]string{
			"spuIdList": spuIdList,
			"saasId":    saasId,
		},
	}
	jsonData, httpErr := c.doRequestV2("goodsOpenApi.v1.spu.detail", params)
	if httpErr != nil {
		return response.CpsSkuDetail{}, httpErr
	}
	var res response.CpsSkuDetail
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

func (c *TencentCloudClient) GetSkuDetailV1(uniqId string) (response.CpsSkuDetail, error) {
	params := map[string]string{
		"request_id": uuid.NewV4().String(),
		"uniq_id":    uniqId,
	}
	jsonData, httpErr := c.doRequestV1("/cps/sku/detail", params, "post")
	if httpErr != nil {
		return response.CpsSkuDetail{}, httpErr
	}
	var res response.CpsSkuDetail
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

func (c *TencentCloudClient) GetSkuListV1(page, pageSize uint, source string) (response.CpsSkuList, error) {
	params := map[string]string{
		"request_id":     uuid.NewV4().String(),
		"page_index":     cast.ToString(page),
		"page_size":      cast.ToString(pageSize),
		"product_source": source,
	}
	jsonData, httpErr := c.doRequestV1("/cps/sku/list", params, "post")
	if httpErr != nil {
		return response.CpsSkuList{}, httpErr
	}
	var res response.CpsSkuList
	err := json.Unmarshal([]byte(jsonData), &res)
	return response.CpsSkuList{}, err
}

func (c *TencentCloudClient) GenerateUrl(spuId, saasId, couponId, txCpsId string) (response.GenerateUrl, error) {
	params := map[string]interface{}{
		"parameter": map[string]string{
			"spuIdList": spuId,
			"saasId":    saasId,
			"sharerId":  c.ShareId,
			"txCpsId":   txCpsId,
			"couponId":  couponId,
		},
	}
	jsonData, httpErr := c.doRequestV2("tradeapi.v1.miniAppOpenApi.generateUrl", params)
	if httpErr != nil {
		return response.GenerateUrl{}, httpErr
	}
	var res response.GenerateUrl
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}

func (c *TencentCloudClient) GenerateUrlNew(param request.GenerateUrl) (response.GenerateUrl, error) {
	param.SharerId = c.ShareId
	params := map[string]interface{}{
		"parameter": param,
	}
	jsonData, httpErr := c.doRequestV2("tradeapi.v1.miniAppOpenApi.generateUrl", params)
	if httpErr != nil {
		return response.GenerateUrl{}, httpErr
	}
	var res response.GenerateUrl
	err := json.Unmarshal([]byte(jsonData), &res)
	return res, err
}
