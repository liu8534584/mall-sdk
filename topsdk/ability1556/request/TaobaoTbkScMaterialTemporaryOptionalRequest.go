package request

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability1556/domain"
	"github.com/liu8534584/mall-sdk/topsdk/util"
)

type TaobaoTbkScMaterialTemporaryOptionalRequest struct {
	/*
	   商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间     */
	StartDsr *int32 `json:"start_dsr,omitempty" required:"false" `
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int32 `json:"page_size,omitempty" required:"false" `
	/*
	   第几页，默认：１ defalutValue��1    */
	PageNo *int32 `json:"page_no,omitempty" required:"false" `
	/*
	   链接形式-1：PC，2：无线，默认为１ defalutValue��1    */
	Platform *int32 `json:"platform,omitempty" required:"false" `
	/*
	   商品筛选-淘客佣金比率上限。如：1234表示12.34%     */
	EndTkRate *int32 `json:"end_tk_rate,omitempty" required:"false" `
	/*
	   商品筛选-淘客佣金比率下限。如：1234表示12.34%     */
	StartTkRate *int32 `json:"start_tk_rate,omitempty" required:"false" `
	/*
	   商品筛选-折扣价范围上限。单位：元     */
	EndPrice *int32 `json:"end_price,omitempty" required:"false" `
	/*
	   商品筛选-折扣价范围上限。单位：元     */
	StartPrice *int32 `json:"start_price,omitempty" required:"false" `
	/*
	   商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限     */
	IsOverseas *bool `json:"is_overseas,omitempty" required:"false" `
	/*
	   商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限     */
	IsTmall *bool `json:"is_tmall,omitempty" required:"false" `
	/*
	   排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）     */
	Sort *string `json:"sort,omitempty" required:"false" `
	/*
	   商品筛选-所在地     */
	Itemloc *string `json:"itemloc,omitempty" required:"false" `
	/*
	   商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到     */
	Cat *string `json:"cat,omitempty" required:"false" `
	/*
	   查询的关键词     */
	Q *string `json:"q,omitempty" required:"false" `
	/*
	   mm_xxx_xxx_xxx的第3段数字     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   mm_xxx_xxx_xxx的第2段数字     */
	SiteId *int64 `json:"site_id" required:"true" `
	/*
	   不传时默认物料id=2836。如果直接对消费者投放，可使用官方个性化算法优化的搜索物料id=17004 defalutValue��2836    */
	MaterialId *int64 `json:"material_id,omitempty" required:"false" `
	/*
	   优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限     */
	HasCoupon *bool `json:"has_coupon,omitempty" required:"false" `
	/*
	   ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
	Ip *string `json:"ip,omitempty" required:"false" `
	/*
	   商品筛选-牛皮癣程度。取值：1不限，2无，3轻微 defalutValue��1    */
	NpxLevel *int32 `json:"npx_level,omitempty" required:"false" `
	/*
	   商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludeRfdRate *bool `json:"include_rfd_rate,omitempty" required:"false" `
	/*
	   商品筛选(特定媒体支持)-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludeGoodRate *bool `json:"include_good_rate,omitempty" required:"false" `
	/*
	   商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludePayRate30 *bool `json:"include_pay_rate_30,omitempty" required:"false" `
	/*
	   商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限     */
	NeedPrepay *bool `json:"need_prepay,omitempty" required:"false" `
	/*
	   商品筛选-是否包邮。true表示包邮，false或不设置表示不限     */
	NeedFreeShipment *bool `json:"need_free_shipment,omitempty" required:"false" `
	/*
	   商品筛选-KA媒体淘客佣金率上限。如：1234表示12.34%     */
	EndKaTkRate *int32 `json:"end_ka_tk_rate,omitempty" required:"false" `
	/*
	   商品筛选-KA媒体淘客佣金率下限。如：1234表示12.34%     */
	StartKaTkRate *int32 `json:"start_ka_tk_rate,omitempty" required:"false" `
	/*
	   智能匹配-设备号加密后的值（MD5加密需32位小写）     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   智能匹配-设备号加密类型：MD5     */
	DeviceEncrypt *string `json:"device_encrypt,omitempty" required:"false" `
	/*
	   智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
	/*
	   商品筛选-锁佣结束时间     */
	LockRateEndTime *int64 `json:"lock_rate_end_time,omitempty" required:"false" `
	/*
	   商品筛选-锁佣开始时间     */
	LockRateStartTime *int64 `json:"lock_rate_start_time,omitempty" required:"false" `
	/*
	   商家筛选-商家id，仅支持饿了么卡券商家id，支持批量请求1-100以内，多个商家id使用英文逗号分隔     */
	SellerIds *string `json:"seller_ids,omitempty" required:"false" `
	/*
	   本地化入参-LBS信息-国标城市码，仅支持单个请求，请求饿了么卡券物料时，该字段必填。 （详细城市ID见：https://mo.m.taobao.com/page_2020010315120200508）     */
	CityCode *string `json:"city_code,omitempty" required:"false" `
	/*
	   本地化入参-LBS信息-纬度     */
	Latitude *string `json:"latitude,omitempty" required:"false" `
	/*
	   本地化入参-LBS信息-经度     */
	Longitude *string `json:"longitude,omitempty" required:"false" `
	/*
	   渠道关系ID，仅适用于渠道推广场景     */
	RelationId *string `json:"relation_id,omitempty" required:"false" `
	/*
	   会员运营ID，仅适用于会员运营场景     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   本地化业务入参-分页唯一标识，非首页的请求必传，值为上一页返回结果中的page_result_key字段值     */
	PageResultKey *string `json:"page_result_key,omitempty" required:"false" `
	/*
	   人群ID，仅适用于物料评估场景material_id=41377     */
	UcrowdId *int64 `json:"ucrowd_id,omitempty" required:"false" `
	/*
	   物料评估-商品列表     */
	UcrowdRankItems *[]domain.TaobaoTbkScMaterialTemporaryOptionalUcrowdrankitems `json:"ucrowd_rank_items,omitempty" required:"false" `
	/*
	   是否获取前N件佣金信息，0否，1是，其他值否     */
	GetTopnRate *int32 `json:"get_topn_rate,omitempty" required:"false" `
}

func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetStartDsr(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.StartDsr = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetPageSize(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetPageNo(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetPlatform(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Platform = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetEndTkRate(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.EndTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetStartTkRate(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.StartTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetEndPrice(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.EndPrice = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetStartPrice(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.StartPrice = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIsOverseas(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.IsOverseas = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIsTmall(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.IsTmall = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetSort(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Sort = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetItemloc(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Itemloc = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetCat(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Cat = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetQ(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Q = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetAdzoneId(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetSiteId(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.SiteId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetMaterialId(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.MaterialId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetHasCoupon(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.HasCoupon = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIp(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Ip = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetNpxLevel(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.NpxLevel = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIncludeRfdRate(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.IncludeRfdRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIncludeGoodRate(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.IncludeGoodRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetIncludePayRate30(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.IncludePayRate30 = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetNeedPrepay(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.NeedPrepay = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetNeedFreeShipment(v bool) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.NeedFreeShipment = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetEndKaTkRate(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.EndKaTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetStartKaTkRate(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.StartKaTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetDeviceValue(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.DeviceValue = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetDeviceEncrypt(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.DeviceEncrypt = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetDeviceType(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.DeviceType = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetLockRateEndTime(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.LockRateEndTime = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetLockRateStartTime(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.LockRateStartTime = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetSellerIds(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.SellerIds = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetCityCode(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.CityCode = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetLatitude(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Latitude = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetLongitude(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.Longitude = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetRelationId(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetSpecialId(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetPageResultKey(v string) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.PageResultKey = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetUcrowdId(v int64) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.UcrowdId = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetUcrowdRankItems(v []domain.TaobaoTbkScMaterialTemporaryOptionalUcrowdrankitems) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.UcrowdRankItems = &v
	return s
}
func (s *TaobaoTbkScMaterialTemporaryOptionalRequest) SetGetTopnRate(v int32) *TaobaoTbkScMaterialTemporaryOptionalRequest {
	s.GetTopnRate = &v
	return s
}

func (req *TaobaoTbkScMaterialTemporaryOptionalRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.StartDsr != nil {
		paramMap["start_dsr"] = *req.StartDsr
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.Platform != nil {
		paramMap["platform"] = *req.Platform
	}
	if req.EndTkRate != nil {
		paramMap["end_tk_rate"] = *req.EndTkRate
	}
	if req.StartTkRate != nil {
		paramMap["start_tk_rate"] = *req.StartTkRate
	}
	if req.EndPrice != nil {
		paramMap["end_price"] = *req.EndPrice
	}
	if req.StartPrice != nil {
		paramMap["start_price"] = *req.StartPrice
	}
	if req.IsOverseas != nil {
		paramMap["is_overseas"] = *req.IsOverseas
	}
	if req.IsTmall != nil {
		paramMap["is_tmall"] = *req.IsTmall
	}
	if req.Sort != nil {
		paramMap["sort"] = *req.Sort
	}
	if req.Itemloc != nil {
		paramMap["itemloc"] = *req.Itemloc
	}
	if req.Cat != nil {
		paramMap["cat"] = *req.Cat
	}
	if req.Q != nil {
		paramMap["q"] = *req.Q
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.SiteId != nil {
		paramMap["site_id"] = *req.SiteId
	}
	if req.MaterialId != nil {
		paramMap["material_id"] = *req.MaterialId
	}
	if req.HasCoupon != nil {
		paramMap["has_coupon"] = *req.HasCoupon
	}
	if req.Ip != nil {
		paramMap["ip"] = *req.Ip
	}
	if req.NpxLevel != nil {
		paramMap["npx_level"] = *req.NpxLevel
	}
	if req.IncludeRfdRate != nil {
		paramMap["include_rfd_rate"] = *req.IncludeRfdRate
	}
	if req.IncludeGoodRate != nil {
		paramMap["include_good_rate"] = *req.IncludeGoodRate
	}
	if req.IncludePayRate30 != nil {
		paramMap["include_pay_rate_30"] = *req.IncludePayRate30
	}
	if req.NeedPrepay != nil {
		paramMap["need_prepay"] = *req.NeedPrepay
	}
	if req.NeedFreeShipment != nil {
		paramMap["need_free_shipment"] = *req.NeedFreeShipment
	}
	if req.EndKaTkRate != nil {
		paramMap["end_ka_tk_rate"] = *req.EndKaTkRate
	}
	if req.StartKaTkRate != nil {
		paramMap["start_ka_tk_rate"] = *req.StartKaTkRate
	}
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceEncrypt != nil {
		paramMap["device_encrypt"] = *req.DeviceEncrypt
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	if req.LockRateEndTime != nil {
		paramMap["lock_rate_end_time"] = *req.LockRateEndTime
	}
	if req.LockRateStartTime != nil {
		paramMap["lock_rate_start_time"] = *req.LockRateStartTime
	}
	if req.SellerIds != nil {
		paramMap["seller_ids"] = *req.SellerIds
	}
	if req.CityCode != nil {
		paramMap["city_code"] = *req.CityCode
	}
	if req.Latitude != nil {
		paramMap["latitude"] = *req.Latitude
	}
	if req.Longitude != nil {
		paramMap["longitude"] = *req.Longitude
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.PageResultKey != nil {
		paramMap["page_result_key"] = *req.PageResultKey
	}
	if req.UcrowdId != nil {
		paramMap["ucrowd_id"] = *req.UcrowdId
	}
	if req.UcrowdRankItems != nil {
		paramMap["ucrowd_rank_items"] = util.ConvertStructList(*req.UcrowdRankItems)
	}
	if req.GetTopnRate != nil {
		paramMap["get_topn_rate"] = *req.GetTopnRate
	}
	return paramMap
}

func (req *TaobaoTbkScMaterialTemporaryOptionalRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}