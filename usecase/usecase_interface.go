package usecase

import "final_project/models"

type UserUsecaseInterface interface {
	GetFollower(username string) models.UserAPIDetail
	GetDetail(userId string) models.UserAPIDetail
}
