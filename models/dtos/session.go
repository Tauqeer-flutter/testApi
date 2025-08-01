package dtos

import "time"

type Session struct {
	Id                uint      `gorm:"primaryKey;autoIncrement" json:"id" required:"false"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	WorkDuration      int       `json:"work_duration"`
	BreakDuration     int       `json:"break_duration"`
	ExtraDuration     int       `json:"extra_duration"`
	SessionType       string    `json:"session_type"`
	WorkStartFilePath string    `json:"work_start_file_path"`
	WorkEndFilePath   string    `json:"work_end_file_path"`
	UserId            uint      `gorm:"Constraint:OnDelete:CASCADE;foreignKey:UserId;references:Id" json:"user_id"`
	Breaks            []Break   `json:"breaks"`
}

type Break struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id" required:"false"`
	Start         time.Time `json:"start"`
	End           time.Time `json:"end"`
	StartFilePath string    `json:"start_file_path"`
	EndFilePath   string    `json:"end_file_path"`
	SessionId     uint      `gorm:"Constrain:OnDelete:CASCADE;foreignKey:SessionId;references:Id" json:"session_id"`
}
