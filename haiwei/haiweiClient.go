package haiwei

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"sort"
	"strings"
	"time"
)

const (
	Kfc      = "kfc"
	Mcdonald = "mdl"
	//必胜客
	Bsk = "bsk"
	//星巴克
	Xbk = "xbk"
	//瑞幸
	Ruixing = "ruixing"
	//365出游
	Chuyou365 = "chuyou365"
	Movie     = "movie"
	//365鲜花
	Xianhua365 = "xianhua365"
	//百果园
	Baiguoyuan = "baiguoyuan"
	// 寄快递
	ExpressDelivery = "expressDelivery"

	//喜茶
	Xicha = "xicha"
	//汉堡王
	HanBaoWang = "hanBaoWang"
	//奈雪的茶
	NaiXueDeCha = "naiXueDeCha"

	More = "more"
)

var ActivityidUrlMap = map[string]string{
	Kfc:             "api/entrance",
	Mcdonald:        "api/mcdonald/entrance",
	Bsk:             "privilege-api/pizzaHut/auth/index",
	Xbk:             "sbkplus/auth/index",
	Ruixing:         "privilege-api/luckin/auth/index",
	Chuyou365:       "privilege-api/tourism/auth/index",
	Movie:           "ticket/auth/index",
	Xianhua365:      "privilege-api/flowerCake/auth/index",
	Baiguoyuan:      "privilege-api/pagoda/auth/index",
	ExpressDelivery: "privilege-api/express/delivery/auth/index",
	More:            "api/index/entrance",
	Xicha:           "privilege-api/heytea/auth/index",
	NaiXueDeCha:     "api/nayuki/entrance",
	HanBaoWang:      "privilege-api/burgerKing/auth/index",
}

type HaiweiClient struct {
	ShareCode  string
	SecretKey  string
	gatewayUrl string
}

func NewHaiweiClient(shareCode, secretKey string) *HaiweiClient {
	return &HaiweiClient{
		ShareCode:  shareCode,
		SecretKey:  secretKey,
		gatewayUrl: "https://ot.jfshou.cn/",
	}
}

func (c *HaiweiClient) GetIndexUrl(pid string) string {
	params := make(map[string]string)
	params["user_id"] = pid
	params["share_code"] = c.ShareCode
	params["timestamp"] = cast.ToString(time.Now().UnixMilli())
	params["extra"] = pid
	params["sign"] = c.generateSign(params)

	return utils.SetQuery("https://ot.jfshou.cn/api/index/entrance", params)
}

func (c *HaiweiClient) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var ss = ""
	for _, v := range keys {
		ss += utils.LineToHump(v) + "=" + params[v] + "&"
	}
	ss += "secretKey=" + c.SecretKey
	h := md5.New()
	h.Write([]byte(ss))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (c *HaiweiClient) GetUrl(ctx context.Context, method, aliPayUserId string, userId string) (string, error) {
	var sysParams = make(map[string]string)
	sysParams["timestamp"] = cast.ToString(time.Now().Unix())
	sysParams["shareCode"] = c.ShareCode
	sysParams["userId"] = aliPayUserId
	sysParams["extra"] = userId
	sysParams["sign"] = c.generateSign(sysParams)

	return utils.SetQuery(c.gatewayUrl+method, sysParams), nil

}

func (c *HaiweiClient) WxGetUrl(ctx context.Context, method, appId, openId string, userId uint64) (string, error) {
	var sysParams = make(map[string]string)
	sysParams["timestamp"] = cast.ToString(time.Now().Unix())
	sysParams["shareCode"] = c.ShareCode
	sysParams["userId"] = cast.ToString(userId)
	sysParams["wxOpenid"] = openId
	sysParams["wxAppid"] = appId
	sysParams["extra"] = cast.ToString(userId)
	sysParams["sign"] = c.generateSign(sysParams)

	return utils.SetQuery(c.gatewayUrl+method, sysParams), nil
}
