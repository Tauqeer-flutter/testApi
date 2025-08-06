package requests

type BreakSessionRequest struct {
	SessionId     *int `json:"sessionId" validate:"required"`
	WorkDuration  *int `json:"workDuration" validate:"required"`
	BreakDuration *int `json:"breakDuration" validate:"required"`
	ExtraDuration *int `json:"extraDuration" validate:"required"`
}
