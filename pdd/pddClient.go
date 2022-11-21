package pdd

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

type pddClient struct {
	ClientId     string
	ClientSecret string
	gatewayUrl   string
}

func NewPddClient(clientId, clientSecret string) *pddClient {
	return &pddClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		gatewayUrl:   "https://gw-api.pinduoduo.com/api/router",
	}
}

func (c *pddClient) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = c.ClientSecret
	for _, v := range keys {
		ss += v + params[v]
	}
	ss += c.ClientSecret
	h := md5.New()
	h.Write([]byte(ss))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (c *pddClient) Execute(req Request, isPost bool) (string, error) {
	var sysParams = make(map[string]string)
	sysParams["client_id"] = c.ClientId
	sysParams["timestamp"] = cast.ToString(time.Now().Unix() / 1000)
	sysParams["data_type"] = "JSON"
	for k, v := range req.GetApiParams() {
		sysParams[k] = cast.ToString(v)
	}
	sysParams["sign"] = c.generateSign(sysParams)

	if isPost {
		return utils.HttpPost(c.gatewayUrl, sysParams)
	}

	response, err := utils.GetResponse(utils.SetQuery(c.gatewayUrl, sysParams), nil)
	if err != nil {
		return "", err
	}
	return string(response), nil

}
