package pddRequest

import "encoding/json"

type PddDdkResourceUrlGen struct {
	typeName         string
	apiParas         map[string]interface{}
	customParameters string
	pid              string
	resourceType     int
	url              string
	generateWeApp    bool
}

func (p *PddDdkResourceUrlGen) GenerateWeApp() bool {
	return p.generateWeApp
}

func (p *PddDdkResourceUrlGen) SetGenerateWeApp(generateWeApp bool) {
	p.apiParas["generate_we_app"] = generateWeApp
	p.generateWeApp = generateWeApp
}

func NewPddDdkResourceUrlGen() *PddDdkResourceUrlGen {
	p := &PddDdkResourceUrlGen{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkResourceUrlGen) Url() string {
	return p.url
}

func (p *PddDdkResourceUrlGen) SetUrl(url string) {
	p.apiParas["url"] = url
	p.url = url
}

func (p *PddDdkResourceUrlGen) ResourceType() int {
	return p.resourceType
}

func (p *PddDdkResourceUrlGen) SetResourceType(resourceType int) {
	p.apiParas["resource_type"] = resourceType
	p.resourceType = resourceType
}

func (p *PddDdkResourceUrlGen) Pid() string {
	return p.pid
}

func (p *PddDdkResourceUrlGen) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *PddDdkResourceUrlGen) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkResourceUrlGen) SetCustomParameters(uid string, sid string) {
	var customParametersData map[string]string
	customParametersData = make(map[string]string)
	customParametersData["uid"] = uid
	customParametersData["sid"] = sid

	customParameters, _ := json.Marshal(customParametersData)
	p.apiParas["custom_parameters"] = string(customParameters)
	p.customParameters = string(customParameters)
}

func (p *PddDdkResourceUrlGen) TypeName() string {
	return p.typeName
}

func (p *PddDdkResourceUrlGen) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func (p *PddDdkResourceUrlGen) GetApiParams() map[string]interface{} {
	return p.apiParas
}

func (p *PddDdkResourceUrlGen) GetApiMethodName() string {
	return "pdd.ddk.resource.url.gen"
}
