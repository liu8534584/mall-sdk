package meituan

type DjRequest interface {
	GetApiParamsObject() interface{}
	GetApiPath() string
	SetConfig(*MeituanDjConfig)
	GetConfig() *MeituanDjConfig
}

type DjBaseRequest struct {
	config *MeituanDjConfig
	client *MeiuanDjClient
}

func (d *DjBaseRequest) GetClient() *MeiuanDjClient {
	return d.client
}

func (d *DjBaseRequest) SetClient(client *MeiuanDjClient) {
	d.client = client
}

func (d *DjBaseRequest) GetConfig() *MeituanDjConfig {
	return d.config
}

func (d *DjBaseRequest) SetConfig(config *MeituanDjConfig) {
	d.config = config
}
