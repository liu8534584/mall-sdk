package pddRequest

import "encoding/json"

type PddDdkGoodsPidGenerate struct {
	typeName    string
	apiParas    map[string]interface{}
	number      int
	pidNameList []string
	mediaId     int64
}

func (p *PddDdkGoodsPidGenerate) MediaId() int64 {
	return p.mediaId
}

func (p *PddDdkGoodsPidGenerate) SetMediaId(mediaId int64) {
	p.apiParas["media_id"] = mediaId
	p.mediaId = mediaId
}

func (p *PddDdkGoodsPidGenerate) PidNameList() []string {
	return p.pidNameList
}

func (p *PddDdkGoodsPidGenerate) SetPidNameList(pidNameList []string) {
	pidStr, _ := json.Marshal(pidNameList)
	p.apiParas["p_id_name_list"] = string(pidStr)
	p.pidNameList = pidNameList
}

func (p *PddDdkGoodsPidGenerate) Number() int {
	return p.number
}

func (p *PddDdkGoodsPidGenerate) SetNumber(number int) {
	p.apiParas["number"] = number
	p.number = number
}

func (p *PddDdkGoodsPidGenerate) ApiParas() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsPidGenerate) TypeName() string {
	return p.typeName
}

func (p *PddDdkGoodsPidGenerate) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkGoodsPidGenerate() *PddDdkGoodsPidGenerate {
	p := &PddDdkGoodsPidGenerate{
		apiParas: make(map[string]interface{}),
	}
	//fmt.Println(p.GetApiMethodName())
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsPidGenerate) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkGoodsPidGenerate) GetApiMethodName() string {
	return "pdd.ddk.goods.pid.generate"
}
