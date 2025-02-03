package models


type UserRequest_DTO struct {
	Task string `json:"task"`
	ScheduledAt string `json:"scheduled_at"`
}