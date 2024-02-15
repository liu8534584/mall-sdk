package meituanCard

type MeituanCardRequest interface {
	GetApiParamsObject() interface{}
	GetApiPath() string
	SetConfig(*MeituanCardConfig)
	GetConfig() *MeituanCardConfig
	IsPost() bool
}

type DdBaseRequest struct {
	config *MeituanCardConfig
	client *MeituanCardClient
}

func (d *DdBaseRequest) GetClient() *MeituanCardClient {
	return d.client
}

func (d *DdBaseRequest) SetClient(client *MeituanCardClient) {
	d.client = client
}

func (d *DdBaseRequest) GetConfig() *MeituanCardConfig {
	return d.config
}

func (d *DdBaseRequest) SetConfig(config *MeituanCardConfig) {
	d.config = config
}
