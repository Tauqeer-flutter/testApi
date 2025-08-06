package requests

import (
	"time"
)

type EndBreakRequest struct {
	SessionId     *uint      `json:"sessionId" form:"sessionId" validate:"required"`
	WorkDuration  *int       `json:"workDuration" form:"workDuration" validate:"required"`
	BreakDuration *int       `json:"breakDuration" form:"breakDuration" validate:"required"`
	ExtraDuration *int       `json:"extraDuration" form:"extraDuration" validate:"required"`
	BreakStart    *time.Time `json:"breakStart" form:"breakStart" validate:"required"`
	BreakEnd      *time.Time `json:"breakEnd" form:"breakEnd" validate:"required"`
}
