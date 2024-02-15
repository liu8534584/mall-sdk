package pddRequest

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

func (p *PddDdkResourceUrlGen) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
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
