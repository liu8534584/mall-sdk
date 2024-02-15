package meituanDd

type MeituanDdRequest interface {
	GetApiParamsObject() interface{}
	GetApiPath() string
	SetConfig(*MeituanDdConfig)
	GetConfig() *MeituanDdConfig
	IsPost() bool
}

type DdBaseRequest struct {
	config *MeituanDdConfig
	client *MeiuanDdClient
}

func (d *DdBaseRequest) GetClient() *MeiuanDdClient {
	return d.client
}

func (d *DdBaseRequest) SetClient(client *MeiuanDdClient) {
	d.client = client
}

func (d *DdBaseRequest) GetConfig() *MeituanDdConfig {
	return d.config
}

func (d *DdBaseRequest) SetConfig(config *MeituanDdConfig) {
	d.config = config
}
