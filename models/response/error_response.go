package response

type BaseResponse struct {
	Status  bool   `json:"status" default:"false"`
	Message string `json:"message"`
}
