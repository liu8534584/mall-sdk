package pddRequest

type PddDdkMemberAuthorityQuery struct {
	typeName         string
	apiParas         map[string]interface{}
	customParameters string
	pid              string
}

func (p *PddDdkMemberAuthorityQuery) Pid() string {
	return p.pid
}

func (p *PddDdkMemberAuthorityQuery) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *PddDdkMemberAuthorityQuery) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkMemberAuthorityQuery) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkMemberAuthorityQuery) TypeName() string {
	return p.typeName
}

func (p *PddDdkMemberAuthorityQuery) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func NewPddDdkMemberAuthorityQuery() *PddDdkMemberAuthorityQuery {
	p := &PddDdkMemberAuthorityQuery{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkMemberAuthorityQuery) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkMemberAuthorityQuery) GetApiMethodName() string {
	return "pdd.ddk.member.authority.query"
}
