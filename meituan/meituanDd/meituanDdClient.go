package meituanDd

import (
	"encoding/hex"
	"encoding/json"
	"github.com/forgoer/openssl"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	DdBaseUrl = "https://union.dianping.com"
)

// MeiuanDdClient 美团到店 https://pub.meituan.com/v2/csr#/apidoc
type MeiuanDdClient struct {
}

func NewMeituanDdClient() *MeiuanDdClient {
	return &MeiuanDdClient{}
}

func (m *MeiuanDdClient) Execute(params MeituanDdRequest) (string, error) {
	marshal, _ := json.Marshal(params.GetApiParamsObject())

	if params.IsPost() {
		return m.HttpPostJson(params.GetApiPath(), string(marshal))
	}
	data := map[string]interface{}{}
	_ = json.Unmarshal(marshal, &data)
	param := map[string]string{}
	for k, v := range data {
		param[k] = cast.ToString(v)
	}

	return utils.HttpGet(utils.SetQuery(params.GetApiPath(), param), time.Second*3)
}

func (m *MeiuanDdClient) HttpPostJson(url string, data string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil && req == nil {
		return "", err
	}

	if req == nil {
		return "", err
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
			ResponseHeaderTimeout: time.Second * 3,
		},
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	if resp == nil {
		return "", err
	}
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (m *MeiuanDdClient) GetAccessToken(appKey, utmSource string, timestamp int64) string {
	src := []byte(utmSource + cast.ToString(timestamp))
	key := []byte(appKey)
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	return hex.EncodeToString(dst)
}

func (m *MeiuanDdClient) GetUtmMedium(appKey, pid string) string {
	src := []byte(pid)
	key := []byte(appKey)
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	return strings.ToUpper(hex.EncodeToString(dst))
}

var DefaultMeituanDdClient *MeiuanDdClient = NewMeituanDdClient()
