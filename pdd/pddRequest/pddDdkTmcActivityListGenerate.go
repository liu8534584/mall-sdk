package pddRequest

type PddDdkTmcActivityListGenerate struct {
	typeName       string
	apiParas       map[string]interface{}
	pageNum        int32
	pageSize       int32
	startTimeLower string
	startTimeUpper string
}

func (p *PddDdkTmcActivityListGenerate) StartTimeUpper() string {
	return p.startTimeUpper
}

func (p *PddDdkTmcActivityListGenerate) SetStartTimeUpper(startTimeUpper string) {
	p.startTimeUpper = startTimeUpper
	p.apiParas["start_time_upper"] = startTimeUpper
}

func (p *PddDdkTmcActivityListGenerate) StartTimeLower() string {
	return p.startTimeLower
}

func (p *PddDdkTmcActivityListGenerate) SetStartTimeLower(startTimeLower string) {
	p.startTimeLower = startTimeLower
	p.apiParas["start_time_lower"] = startTimeLower
}

func (p *PddDdkTmcActivityListGenerate) PageNum() int32 {
	return p.pageNum
}

func (p *PddDdkTmcActivityListGenerate) SetPageNum(pageNum int32) {
	p.pageNum = pageNum
	p.apiParas["page_num"] = pageNum
}

func (p *PddDdkTmcActivityListGenerate) PageSize() int32 {
	return p.pageSize
}

func (p *PddDdkTmcActivityListGenerate) SetPageSize(pageSize int32) {
	p.pageSize = pageSize
	p.apiParas["page_size"] = pageSize
}

func (p *PddDdkTmcActivityListGenerate) TypeName() string {
	return p.typeName
}

func (p *PddDdkTmcActivityListGenerate) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkTmcActivityListGenerate() *PddDdkTmcActivityListGenerate {
	p := &PddDdkTmcActivityListGenerate{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkTmcActivityListGenerate) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkTmcActivityListGenerate) GetApiMethodName() string {
	return "pdd.ddk.tmc.activity.list"
}
