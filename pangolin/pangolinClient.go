package pangolin

import (
	"context"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Client struct {
	appId       string
	securityKey string
	gatewayUrl  string
}

func NewClient(appId string, securityKey string) *Client {
	return &Client{
		appId:       appId,
		securityKey: securityKey,
		gatewayUrl:  "https://ecom.pangolin-sdk-toutiao.com",
	}
}

type ReqBody struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version"`
	Sign      string `json:"sign"`
	SignType  string `json:"sign_type"`
	ReqId     string `json:"req_id"`
	Data      string `json:"data"`
}

func (c *Client) execute(ctx context.Context, method string, data interface{}) (string, error) {

	dataBytes, _ := json.Marshal(data)

	req := ReqBody{
		AppId:     c.appId,
		Timestamp: time.Now().Unix(),
		SignType:  "MD5",
		ReqId:     uuid.NewV4().String(),
		Version:   "1",
		Data:      string(dataBytes),
	}

	sign := getSign(req, c.securityKey)

	req.Sign = sign

	postDataBytes, _ := json.Marshal(req)
	resp, _, err := PangolinPost(c.gatewayUrl+method, string(postDataBytes))
	if err != nil {
		return "", err
	}

	return resp, nil

}

type ProductLinkRequest struct {
	ProductUrl   string `json:"product_url,omitempty"`   // 是  https://www.a.com/products/1 商品url。与商品接口detail_url一致
	ProductExt   string `json:"product_ext,omitempty"`   //否 sdfdjfdlsfj 商品搜索接口返回的product.ext 字段, 尽量填写
	ExternalInfo string `json:"external_info,omitempty"` // 否 媒体传递扩展参数的字段， 字符只允许字母大小写、数字、下划线，长度不超过20
	ShareType    []int  `json:"share_type,omitempty"`    // 否 [1,2,3] 转链类型：1、抖音 deep link；2、抖音二维码；3、抖音口令。 不填默认只有1
}

type ProductLinkResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data struct {
		DyPassword string `json:"dy_password"` //抖音口令。当share_type包含3时会返回
		DyQrCode   struct {
			Url    string `json:"url"`    //二维码地址
			Width  int    `json:"width"`  //二维码宽度
			Height int    `json:"height"` // 二维码高度
		} `json:"dy_qr_code"` //抖音二维码。当share_type包含2时会返回
		DyDeeplink string `json:"dy_deeplink"` //抖音deep link。当share_type包含1时会返回
		DyZlink    string `json:"dy_zlink"`
	} `json:"data"`
}

// 商品转链
func (c *Client) ProductLink(ctx context.Context, productLink ProductLinkRequest) (ProductLinkResponse, error) {
	var response ProductLinkResponse
	executeResp, err := c.execute(ctx, "/product/link", productLink)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(executeResp), &response)
	if err != nil {
		return response, err
	}

	if response.Code != 0 {
		return response, errors.New("该商品返现活动已暂停\n请您更换其他商品5")
	}

	return response, err
}

type AggregateH5Request struct {
	ProductUrl   string `json:"product_url,omitempty"`   // 用于强插指定商品到聚合页头部位置
	MaterialId   string `json:"material_id,omitempty"`   //否 聚合页类型枚举值
	ExternalInfo string `json:"external_info,omitempty"` // 否 媒体传递扩展参数的字段， 字符只允许字母大小写、数字、下划线，长度不超过20
}

type AggregateH5Response struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data struct {
		H5PageLink string `json:"h5_page_link"` //聚合页url链接
		FullScreen struct {
			ShortUrl string `json:"short_url"` //抖音短链
			Qrcode   struct {
				Url    string `json:"url"`    //二维码地址
				Width  int    `json:"width"`  //二维码宽度
				Height int    `json:"height"` // 二维码高度
			} `json:"qrcode"` //抖音二维码
			ShareCommand string `json:"share_command"` //抖音口令
			Deeplink     string `json:"deeplink"`      //抖音deeplink
			DyZlink      string `json:"zlink"`
		} `json:"full_screen"` //跳转抖音全屏商品聚合页的转链信息
		SevenSplitScreen struct {
			ShortUrl string `json:"short_url"` //抖音短链
			Qrcode   struct {
				Url    string `json:"url"`    //二维码地址
				Width  int    `json:"width"`  //二维码宽度
				Height int    `json:"height"` // 二维码高度
			} `json:"qrcode"` //抖音二维码
			ShareCommand string `json:"share_command"` //抖音口令
			Deeplink     string `json:"deeplink"`      //抖音deeplink
		} `json:"seven_split_screen"` //跳转抖音七分屏商品聚合页的转链信息
	} `json:"data"`
}

// 活动转链
func (c *Client) AggregateH5(ctx context.Context, aggregateH5 AggregateH5Request) (AggregateH5Response, error) {
	var response AggregateH5Response
	executeResp, err := c.execute(ctx, "/aggregate/h5", aggregateH5)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal([]byte(executeResp), &response)
	if err != nil {
		return response, err
	}

	if response.Code != 0 {
		return response, errors.New("该返现活动已暂停\n请您更换其他商品6")
	}

	return response, err
}

type OrderListRequest struct {
	Size      int64    `json:"size"` //取值区间： [1, 50]
	Cursor    string   `json:"cursor"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
	OrderIds  []string `json:"order_ids"`
	OrderType int      `json:"order_type"` //1-商品分销订单，2-直播间分销订单
	TimeType  string   `json:"time_type"`  //pay-按支付时间查询 update-按更新时间查询
}

type OrderListResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data struct {
		Cursor string `json:"cursor"`
		Orders []struct {
			OrderId                 string `json:"order_id"`
			AppId                   string `json:"app_id"`
			ProductId               string `json:"product_id"`
			ProductName             string `json:"product_name"`
			AuthorAccount           string `json:"author_account"`
			AdsAttribution          string `json:"ads_attribution"`
			ProductImg              string `json:"product_img"`
			TotalPayAmount          int    `json:"total_pay_amount"`
			PaySuccessTime          string `json:"pay_success_time"`
			RefundTime              string `json:"refund_time"`
			PayGoodsAmount          int    `json:"pay_goods_amount"`
			EstimatedCommission     int    `json:"estimated_commission"`
			SplitRate               int    `json:"split_rate"`
			AfterSalesStatus        int    `json:"after_sales_status"`
			FlowPoint               string `json:"flow_point"`
			ExternalInfo            string `json:"external_info"`
			SettleTime              string `json:"settle_time"`
			ConfirmTime             string `json:"confirm_time"`
			UpdateTime              string `json:"update_time"`
			EstimatedTechServiceFee int    `json:"estimated_tech_service_fee"`
		} `json:"orders"`
	} `json:"data"`
}
