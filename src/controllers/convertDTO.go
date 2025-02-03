package controllers

import (
	"src/models"
	"time"
)

func convertDTO(request_dto models.UserRequest_DTO) (models.UserRequest, error) {
	var request models.UserRequest
	scheduled_at, err := time.Parse(time.RFC3339, request_dto.ScheduledAt)
	if err != nil {
		return models.UserRequest{}, err
	}

	request.Task = request_dto.Task
	request.ScheduledAt = scheduled_at
	

	return request, nil
}