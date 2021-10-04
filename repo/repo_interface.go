package repo

import "final_project/models"

// 3
type UserRepoInterface interface {
	// ConsumeApi AddData list all method function in package repo
	ConsumeApi() models.UserAPI
}
