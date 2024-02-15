package pddRequest

type PddDdkGoodsZsUnitUrlGen struct {
	typeName          string
	apiParas          map[string]interface{}
	customParameters  string
	sourceUrl         string
	pid               string
	generateShortLink bool
}

func (p *PddDdkGoodsZsUnitUrlGen) GenerateShortLink() bool {
	return p.generateShortLink
}

func (p *PddDdkGoodsZsUnitUrlGen) SetGenerateShortLink(generateShortLink bool) {
	p.apiParas["generate_short_link"] = generateShortLink
	p.generateShortLink = generateShortLink
}

func (p *PddDdkGoodsZsUnitUrlGen) Pid() string {
	return p.pid
}

func (p *PddDdkGoodsZsUnitUrlGen) SetPid(pid string) {
	p.apiParas["pid"] = pid
	p.pid = pid
}

func (p *PddDdkGoodsZsUnitUrlGen) SourceUrl() string {
	return p.sourceUrl
}

func (p *PddDdkGoodsZsUnitUrlGen) SetSourceUrl(sourceUrl string) {
	p.apiParas["source_url"] = sourceUrl
	p.sourceUrl = sourceUrl
}

func (p *PddDdkGoodsZsUnitUrlGen) CustomParameters() string {
	return p.customParameters
}

func (p *PddDdkGoodsZsUnitUrlGen) SetCustomParameters(customParameters string) {
	p.apiParas["custom_parameters"] = customParameters
	p.customParameters = customParameters
}

func (p *PddDdkGoodsZsUnitUrlGen) TypeName() string {
	return p.typeName
}

func NewPddDdkGoodsZsUnitUrlGen() *PddDdkGoodsZsUnitUrlGen {
	p := &PddDdkGoodsZsUnitUrlGen{
		apiParas: make(map[string]interface{}),
	}
	p.SetTypeName(p.GetApiMethodName())
	return p
}

func (p *PddDdkGoodsZsUnitUrlGen) SetTypeName(typeName string) {
	p.apiParas["type"] = typeName
	p.typeName = typeName
}

func (p *PddDdkGoodsZsUnitUrlGen) GetApiMethodName() string {
	return "pdd.ddk.goods.zs.unit.url.gen"
}

func (p *PddDdkGoodsZsUnitUrlGen) GetApiParams() map[string]interface{} {
	return p.apiParas
}
