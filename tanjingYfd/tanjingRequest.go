package tanjingYfd

type TanjingRequest interface {
	GetApiPath() string
	GetApiParamsObject() interface{}
	GetClient() *TanjingClient
	SetClient(*TanjingClient)
	GetConfig() *TanjingBaseConfig
	SetConfig(*TanjingBaseConfig)
}

type TanjingBaseRequest struct {
	client *TanjingClient
	config *TanjingBaseConfig
}

func (t *TanjingBaseRequest) GetClient() *TanjingClient {
	return t.client
}

func (t *TanjingBaseRequest) SetClient(client *TanjingClient) {
	t.client = client
}

func (t *TanjingBaseRequest) GetConfig() *TanjingBaseConfig {
	return t.config
}

func (t *TanjingBaseRequest) SetConfig(config *TanjingBaseConfig) {
	t.config = config
}
