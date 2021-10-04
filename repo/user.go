package repo

import (
	"encoding/json"
	"final_project/config/env"
	"final_project/models"
	"github.com/go-resty/resty/v2"
)

type UserRepoStruct struct {

}

func NewUserRepoImpl() UserRepoInterface {
	return &UserRepoStruct{}
}


func (u *UserRepoStruct) AddData(usr models.User) {

}

func (u *UserRepoStruct) ConsumeApi() models.UserAPI {

	client := resty.New()
	url := env.Config.UrlJsonUser

	res := models.UserAPI{}
	resp, err := client.R().
		SetResult(res).
		Get(url)

	if err != nil {
		return models.UserAPI{}
	}

	result := resp.String()
	if errs := json.Unmarshal([]byte(result), &res); errs != nil {
		return models.UserAPI{}
	}

	return res
}