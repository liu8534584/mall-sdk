package request

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/jd"
	"github.com/liu8534584/mall-sdk/jd/jdUnionOpenGoodsQuery/response"
)

type UnionOpenGoodsQueryRequest struct {
	jd.BaseJdApiRequest
	Param *UnionOpenGoodsQueryParam
}

func (j *UnionOpenGoodsQueryRequest) GetMethodName() string {
	return "jd.union.open.goods.query"
}

func New(config *jd.JdBaseConfig) *UnionOpenGoodsQueryRequest {
	request := UnionOpenGoodsQueryRequest{
		Param: &UnionOpenGoodsQueryParam{},
	}
	request.SetConfig(config)
	request.SetClient(jd.DefaultJdApiClient)

	return &request
}

func (j *UnionOpenGoodsQueryRequest) Execute() (*response.UnionOpenGoodsQueryResponseResult, error) {
	responseJson, err := j.GetClient().Request(j, false)

	if err != nil {
		return nil, err
	}
	NewResponseObj := &response.UnionOpenGoodsQueryResponseResult{}
	responseObj := map[string]response.UnionOpenGoodsQueryResponse{}
	_ = json.Unmarshal([]byte(responseJson), &responseObj)
	for _, v := range responseObj {
		if v.GetResult != "" {
			_ = json.Unmarshal([]byte(v.GetResult), &NewResponseObj)
		}
		if v.QueryResult != "" {
			_ = json.Unmarshal([]byte(v.QueryResult), &NewResponseObj)
		}
		break
	}

	return NewResponseObj, err
}

func (j *UnionOpenGoodsQueryRequest) GetParamsObject() interface{} {
	return j.Param
}

func (j *UnionOpenGoodsQueryRequest) GetParams() *UnionOpenGoodsQueryParam {
	return j.Param
}

type UnionOpenGoodsQueryParam struct {
	GoodsReqDto GoodsReqDto `json:"goodsReqDTO"`
}

type GoodsReqDto struct {
	PageIndex   int    `json:"pageIndex"`
	ChannelId   int    `json:"channelId"`
	Fields      string `json:"fields"`
	PageSize    int    `json:"pageSize"`         //每页数量，单页数最大30，默认20
	SkuIds      []int  `json:"skuIds,omitempty"` //skuid集合(一次最多支持查询20个sku)，数组类型开发时记得加[]
	Keyword     string `json:"keyword"`
	SortName    string `json:"sortName"`  //排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort        string `json:"sort"`      //	asc,desc升降序,默认降序
	IsCoupon    int    `json:"isCoupon"`  //	是否是优惠券商品，1：有优惠券
	IsPg        int    `json:"isPG"`      //	是否是拼购商品，1：拼购商品
	CouponUrl   string `json:"couponUrl"` //  优惠券链接
	Pid         string `json:"pid"`
	Cid1        int    `json:"cid1,omitempty"`
	ForbidTypes string `json:"forbidTypes"` //10微信京东购物小程序禁售，11微信京喜小程序禁售
	Owner       string `json:"owner"`       //商品类型：自营[g]，POP[p]
}
