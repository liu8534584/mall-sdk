package request

type CommandParseAndShareRequest struct {
	Params *CommandParseAndShareParams
}

func NewCommandParseAndShareRequest() *CommandParseAndShareRequest {
	return &CommandParseAndShareRequest{}
}

type CommandParseAndShareParams struct {
	Uid         int64       `json:"uid"`
	Command     string      `json:"command"`
	ShareParams ShareParams `json:"share_params"`
}

type ShareParams struct {
	Pid           string `json:"pid,omitempty"`
	ExternalInfo  string `json:"external_info,omitempty"`
	Platform      int    `json:"platform,omitempty"`
	NeedQrCode    bool   `json:"need_qr_code,omitempty"`
	NeedShareLink bool   `json:"need_share_link,omitempty"`
	NeedZlink     bool   `json:"need_zlink,omitempty"`
	NeedCommand   bool   `json:"need_command,omitempty"`
}

func (r *CommandParseAndShareRequest) GetParams() interface{} {
	return r.Params
}

func (r *CommandParseAndShareRequest) GetUrlPath() string {
	return "/api/life/v1/outside_distribution/command_parse_and_share/"
}
