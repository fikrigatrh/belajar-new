package controller

import (
	"final_project/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	usr usecase.UserUsecaseInterface
}

func CreateUserController(r *gin.RouterGroup, usr usecase.UserUsecaseInterface)  {
	handler := UserController{usr: usr}

	// create endpoint
	r.GET("/follower/:username", handler.GetFollower)
	r.GET("/:userId/detail", handler.GetDetail)
}

//func (u UserController) AddData(c *gin.Context)  {
//	var usr models.User
//
//	err := c.ShouldBindJSON(&usr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, "error when binding")
//		return
//	}
//
//	log.Println(usr.Nama, "<<<< nama")
//	log.Println(usr.Alamat, "<<< alamat")
//
//	var alamat models.AlamatDetail
//
//	m := usr.Alamat.(map[string]interface{})
//	alamat.Kelurahan = m["kelurahan"].(string)
//
//
//	log.Println(alamat.Kelurahan, "><><><><> kelurahan")
//	 user := u.usr.Hmmm(usr)
//	c.JSON(http.StatusOK, user)
//}

func (u UserController) GetFollower(c *gin.Context)  {
	username := c.Param("username")
	res := u.usr.GetFollower(username)

	c.JSON(http.StatusOK, res)
}

func (u UserController) GetDetail(c *gin.Context) {
	userId := c.Param("userId")
	res := u.usr.GetDetail(userId)

	c.JSON(http.StatusOK, res)
}