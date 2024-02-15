package response

type PidCreateResponse struct {
	ApiResponse
	Data PidCreateData `json:"data"`
}

type PidCreateData struct {
	Pid string `json:"pid"`
}
