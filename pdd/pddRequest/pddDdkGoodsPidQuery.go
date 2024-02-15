package pddRequest

import "encoding/json"

type PddDdkGoodsPidQuery struct {
	typeName string
	apiParas map[string]interface{}
	page     uint64
	pageSize uint64
	pidList  []string
	status   int
}

func (p *PddDdkGoodsPidQuery) Status() int {
	return p.status
}

func (p *PddDdkGoodsPidQuery) SetStatus(status int) {
	p.apiParas["status"] = status
	p.status = status
}

func (p *PddDdkGoodsPidQuery) PidList() []string {
	return p.pidList
}

func (p *PddDdkGoodsPidQuery) SetPidList(pidList []string) {
	pidStr, _ := json.Marshal(pidList)
	p.apiParas["pid_list"] = string(pidStr)
	p.pidList = pidList
}

func (p *PddDdkGoodsPidQuery) PageSize() uint64 {
	return p.pageSize
}

func (p *PddDdkGoodsPidQuery) SetPageSize(pageSize uint64) {
	p.apiParas["page_size"] = pageSize
	p.pageSize = pageSize
}

func NewPddDdkGoodsPidQuery() *PddDdkGoodsPidQuery {
	p := &PddDdkGoodsPidQuery{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsPidQuery) Page() uint64 {
	return p.page
}

func (p *PddDdkGoodsPidQuery) SetPage(page uint64) {
	p.apiParas["page"] = page
	p.page = page
}

func (p *PddDdkGoodsPidQuery) TypeName() string {
	return p.typeName
}

func (p *PddDdkGoodsPidQuery) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func (p *PddDdkGoodsPidQuery) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsPidQuery) GetApiMethodName() string {
	return "pdd.ddk.goods.pid.query"
}
