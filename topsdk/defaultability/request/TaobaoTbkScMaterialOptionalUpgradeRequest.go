package request

type TaobaoTbkScMaterialOptionalUpgradeRequest struct {
	/*
	   商品筛选-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间     */
	StartDsr *int64 `json:"start_dsr,omitempty" required:"false" `
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   第几页，默认：１ defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   商品筛选-淘客佣金比率上限。如：1234表示12.34%     */
	EndTkRate *int64 `json:"end_tk_rate,omitempty" required:"false" `
	/*
	   商品筛选-淘客佣金比率下限。如：1234表示12.34%     */
	StartTkRate *int64 `json:"start_tk_rate,omitempty" required:"false" `
	/*
	   商品筛选-折扣价范围上限。单位：元     */
	EndPrice *int64 `json:"end_price,omitempty" required:"false" `
	/*
	   商品筛选-折扣价范围上限。单位：元     */
	StartPrice *int64 `json:"start_price,omitempty" required:"false" `
	/*
	   商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限     */
	IsOverseas *bool `json:"is_overseas,omitempty" required:"false" `
	/*
	   商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限     */
	IsTmall *bool `json:"is_tmall,omitempty" required:"false" `
	/*
	   排序_des（降序），排序_asc（升序），销量（total_sales），淘客收入比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price），匹配分（match）     */
	Sort *string `json:"sort,omitempty" required:"false" `
	/*
	   商品筛选-所在地     */
	Itemloc *string `json:"itemloc,omitempty" required:"false" `
	/*
	   商品筛选-后台类目ID。用,分割，最大10个     */
	Cat *string `json:"cat,omitempty" required:"false" `
	/*
	   商品筛选-查询词；注意：使用标题精准搜索时，若无消费者比价场景ID2权限，当搜索结果只有一个商品时则出参不再提供商品推广链接和商品id字段，若搜索结果仍有多个商品，则正常出参。同时无消费者比价场景ID2权限，q参数也不再支持入参淘宝复制的商品链接进行搜索查询，仅支持入参含新商品id的淘宝客推广链接如uland链接进行搜索查询(场景id使用说明参考《淘宝客新商品ID升级》白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud)     */
	Q *string `json:"q,omitempty" required:"false" `
	/*
	   推广位id，mm_xxx_xxx_12345678三段式的最后一段数字（登录pub.alimama.com推广管理-推广位管理中查询）     */
	AdzoneId *int64 `json:"adzone_id" required:"true" `
	/*
	   mm_xxx_xxx_xxx的第2段数字     */
	SiteId *int64 `json:"site_id" required:"true" `
	/*
	   物料id，不传时默认物料material_id=80309；如果直接对消费者投放，可使用官方个性化算法优化的搜索物料material_id=17004（注意：若物料id=17004没查询到结果则出系统默认物料id=80309的查询结果） defalutValue��80309    */
	MaterialId *int64 `json:"material_id,omitempty" required:"false" `
	/*
	   优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限     */
	HasCoupon *bool `json:"has_coupon,omitempty" required:"false" `
	/*
	   ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
	Ip *string `json:"ip,omitempty" required:"false" `
	/*
	   商品筛选-牛皮癣程度。取值：1不限，2无，3轻微 defalutValue��1    */
	NpxLevel *int64 `json:"npx_level,omitempty" required:"false" `
	/*
	   商品筛选-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludeRfdRate *bool `json:"include_rfd_rate,omitempty" required:"false" `
	/*
	   商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludeGoodRate *bool `json:"include_good_rate,omitempty" required:"false" `
	/*
	   商品筛选-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限     */
	IncludePayRate30 *bool `json:"include_pay_rate_30,omitempty" required:"false" `
	/*
	   商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限     */
	NeedPrepay *bool `json:"need_prepay,omitempty" required:"false" `
	/*
	   商品筛选-是否包邮。true表示包邮，false或不设置表示不限     */
	NeedFreeShipment *bool `json:"need_free_shipment,omitempty" required:"false" `
	/*
	   智能匹配-设备号加密后的值（MD5加密需32位小写）；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   智能匹配-设备号加密类型：MD5；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
	DeviceEncrypt *string `json:"device_encrypt,omitempty" required:"false" `
	/*
	   智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID；使用智能推荐请先签署协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
	/*
	   渠道关系ID，仅适用于渠道推广场景     */
	RelationId *string `json:"relation_id,omitempty" required:"false" `
	/*
	   会员运营ID，仅适用于会员运营场景     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   是否获取前N件佣金信息，0否，1是，其他值否     */
	GetTopnRate *int64 `json:"get_topn_rate,omitempty" required:"false" `
	/*
	   1-动态ID转链场景，2-消费者比价场景（不填默认为1）；场景id使用说明参考《淘宝客新商品ID升级》白皮书：https://www.yuque.com/taobaolianmengguanfangxiaoer/zmig94/tfyt0pahmlpzu2ud     */
	BizSceneId *string `json:"biz_scene_id,omitempty" required:"false" `
	/*
	   1-自购省，2-推广赚（代理模式专属ID，代理模式必填，非代理模式不用填写该字段）     */
	PromotionType *string `json:"promotion_type,omitempty" required:"false" `
	/*
	   线报内容筛选—内容生产开始时间，13毫秒时间戳     */
	MgcStartTime *string `json:"mgc_start_time,omitempty" required:"false" `
	/*
	   线报内容筛选—内容生产截止时间，13毫秒时间戳     */
	MgcEndTime *string `json:"mgc_end_time,omitempty" required:"false" `
	/*
	   线报状态筛选，0-全部 1-过期 2-实时生效 3-未来生效 不传默认过滤有效     */
	MgcStatus *string `json:"mgc_status,omitempty" required:"false" `
}

func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetStartDsr(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.StartDsr = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetPageSize(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetPageNo(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetEndTkRate(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.EndTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetStartTkRate(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.StartTkRate = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetEndPrice(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.EndPrice = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetStartPrice(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.StartPrice = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIsOverseas(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.IsOverseas = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIsTmall(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.IsTmall = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetSort(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.Sort = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetItemloc(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.Itemloc = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetCat(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.Cat = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetQ(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.Q = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetAdzoneId(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.AdzoneId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetSiteId(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.SiteId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetMaterialId(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.MaterialId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetHasCoupon(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.HasCoupon = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIp(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.Ip = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetNpxLevel(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.NpxLevel = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIncludeRfdRate(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.IncludeRfdRate = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIncludeGoodRate(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.IncludeGoodRate = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetIncludePayRate30(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.IncludePayRate30 = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetNeedPrepay(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.NeedPrepay = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetNeedFreeShipment(v bool) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.NeedFreeShipment = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetDeviceValue(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.DeviceValue = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetDeviceEncrypt(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.DeviceEncrypt = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetDeviceType(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.DeviceType = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetRelationId(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetSpecialId(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetGetTopnRate(v int64) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.GetTopnRate = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetBizSceneId(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.BizSceneId = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetPromotionType(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.PromotionType = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetMgcStartTime(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.MgcStartTime = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetMgcEndTime(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.MgcEndTime = &v
	return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeRequest) SetMgcStatus(v string) *TaobaoTbkScMaterialOptionalUpgradeRequest {
	s.MgcStatus = &v
	return s
}

func (req *TaobaoTbkScMaterialOptionalUpgradeRequest) ToMap() map[string]interface{} {
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
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceEncrypt != nil {
		paramMap["device_encrypt"] = *req.DeviceEncrypt
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.GetTopnRate != nil {
		paramMap["get_topn_rate"] = *req.GetTopnRate
	}
	if req.BizSceneId != nil {
		paramMap["biz_scene_id"] = *req.BizSceneId
	}
	if req.PromotionType != nil {
		paramMap["promotion_type"] = *req.PromotionType
	}
	if req.MgcStartTime != nil {
		paramMap["mgc_start_time"] = *req.MgcStartTime
	}
	if req.MgcEndTime != nil {
		paramMap["mgc_end_time"] = *req.MgcEndTime
	}
	if req.MgcStatus != nil {
		paramMap["mgc_status"] = *req.MgcStatus
	}
	return paramMap
}

func (req *TaobaoTbkScMaterialOptionalUpgradeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
