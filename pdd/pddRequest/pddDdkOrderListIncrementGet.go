package pddRequest

// PddDdkOrderListIncrementGetRequest pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
//https://open.pinduoduo.com/application/document/api?id=pdd.ddk.order.list.increment.get&permissionId=2
type PddDdkOrderListIncrementGet struct {
	typeName        string
	apiParas        map[string]interface{}
	endUpdateTime   int64
	startUpdateTime int64
	page            int
	pageSize        int
	returnCount     bool
}

func NewPddDdkOrderListIncrementGet() *PddDdkOrderListIncrementGet {
	p := &PddDdkOrderListIncrementGet{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkOrderListIncrementGet) SetEndUpdateTime(endUpdateTime int64) {
	p.endUpdateTime = endUpdateTime
	p.apiParas["end_update_time"] = endUpdateTime
}

func (p *PddDdkOrderListIncrementGet) SetStartUpdateTime(startUpdateTime int64) {
	p.startUpdateTime = startUpdateTime
	p.apiParas["start_update_time"] = startUpdateTime
}

func (p *PddDdkOrderListIncrementGet) SetPage(page int) {
	p.page = page
	p.apiParas["page"] = page
}

func (p *PddDdkOrderListIncrementGet) SetPageSize(pageSize int) {
	p.pageSize = pageSize
	p.apiParas["page_size"] = pageSize
}

func (p *PddDdkOrderListIncrementGet) SetReturnCount(returnCount bool) {
	p.returnCount = returnCount
	p.apiParas["return_count"] = returnCount
}

func (p *PddDdkOrderListIncrementGet) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

// GetApiMethodName 返回接口名称
func (p *PddDdkOrderListIncrementGet) GetApiMethodName() string {
	return "pdd.ddk.order.list.increment.get"
}

// GetApiParams 返回请求参数
func (p *PddDdkOrderListIncrementGet) GetApiParams() map[string]interface{} {
	return p.apiParas
}
