package caocao

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"io"
	"sort"
	"time"
)

type CaocaoClient struct {
	ClientId   string
	SecretKey  string
	gatewayUrl string
}

func NewCaocaoClient(clientId, secretKey string) *CaocaoClient {
	return &CaocaoClient{
		ClientId:   clientId,
		SecretKey:  secretKey,
		gatewayUrl: "https://mobile.caocaokeji.cn",
	}
}

func (c *CaocaoClient) GetIndexUrl(pid, phone string, phoneType int) string {
	params := make(map[string]string)
	params["ext_user_id"] = pid
	params["sign"] = c.generateSign(params)
	params["timestamp"] = cast.ToString(time.Now().UnixMilli())
	params["phone"] = phone
	params["phoneType"] = cast.ToString(phoneType)

	return utils.SetQuery(c.gatewayUrl+"/pay-travel/home", params)
}

func (c *CaocaoClient) generateSign(params map[string]string) string {
	params["sign_key"] = c.SecretKey
	params["client_id"] = c.ClientId
	params["timestamp"] = cast.ToString(time.Now().Unix() * 1000)

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = ""
	for _, v := range keys {
		ss += v + params[v]
	}
	h := sha1.New()
	_, _ = io.WriteString(h, ss)
	return hex.EncodeToString(h.Sum([]byte("")))
}
