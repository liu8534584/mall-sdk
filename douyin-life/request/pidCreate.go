package request

type PidCreateRequest struct {
	Params *PidCreateParams
}

func (r *PidCreateRequest) GetParams() interface{} {
	return r.Params
}

func (r *PidCreateRequest) GetUrlPath() string {
	return "/api/life/v1/outside_distribution/create_pid/"
}

type PidCreateParams struct {
	UID       int64  `json:"uid,omitempty"`
	MediaType int    `json:"media_type,omitempty"`
	MediaID   int    `json:"media_id,omitempty"`
	MediaName string `json:"media_name,omitempty"`
	SiteID    int    `json:"site_id,omitempty"`
	SiteName  string `json:"site_name,omitempty"`
}
