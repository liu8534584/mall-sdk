package request

type GenerateUrl struct {
	SaasId        int    `json:"saasId"` //商家id
	SpuId         int    `json:"spuId"`
	SharerId      string `json:"sharerId"`      //分销机构id
	CouponId      int    `json:"couponId"`      //优惠券id
	Url           string `json:"url"`           //需要转链的链接
	TxCpsId       string `json:"txCpsId"`       //自定义推广参数： 分销机构 CPS追踪参数，供机构内部分佣使用，每个自定义推广参数需要保证唯一性；最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~
	ProductSource int    `json:"productSource"` //商品标识：1=YMALL，代表云mall商品；2=OWN，代表第三方商品空：全部
	SkuId         string `json:"skuId"`         //skuId，与warehouseId同时出现
	WarehouseId   string `json:"warehouseId"`   //商品数据源信息，与skuId同时出现，productSource=2 时必填
	MaterialType  string `json:"materialType"`  //转链物料类型(0:商品【默认】,1:活动。与sku_id同时出现)
}
