package usecase

import "final_project/models"

type UserUsecaseInterface interface {
	Hmmm(usr models.User) models.User
	GetFollower(username string) models.UserAPIDetail
	GetDetail(userId string) models.UserAPIDetail
}
