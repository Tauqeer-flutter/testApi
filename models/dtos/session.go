package dtos

import "time"

type Session struct {
	Id                uint       `gorm:"primaryKey;autoIncrement" json:"id" required:"false"`
	StartTime         time.Time  `json:"startTime" validate:"required"`
	EndTime           *time.Time `json:"endTime"`
	Mode              string     `gorm:"NOT NULL" json:"mode" validate:"required"`
	WorkDuration      int        `json:"workDuration"`
	BreakDuration     int        `json:"breakDuration"`
	ExtraDuration     int        `json:"extraDuration"`
	WorkStartFilePath string     `json:"workStartFilePath"`
	WorkEndFilePath   string     `json:"workEndFilePath"`
	UserId            uint       `gorm:"Constraint:OnDelete:CASCADE;foreignKey:UserId;references:Id" json:"userId"`
	Breaks            []Break    `json:"breaks"`
}

type Break struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id" required:"false"`
	Start         time.Time `json:"start"`
	End           time.Time `json:"end"`
	StartFilePath string    `json:"startFilePath"`
	EndFilePath   string    `json:"endFilePath"`
	SessionId     uint      `gorm:"Constrain:OnDelete:CASCADE;foreignKey:SessionId;references:Id" json:"sessionId"`
}
