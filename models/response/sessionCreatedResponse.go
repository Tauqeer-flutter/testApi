package response

import "testApi/models/dtos"

type SessionCreatedResponse struct {
	Status  bool         `json:"status" default:"true"`
	Message string       `json:"message" required:"true"`
	Session dtos.Session `json:"session" required:"true"`
}
