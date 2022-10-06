package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func SaveMoney(c *gin.Context) {

	user := &model.User{}

	err := c.Bind(user)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "save money bind error")
		return
	}

	if user.Id == "" {
		log.Println("id가 비어있습니다")
		module.Response(c, http.StatusBadRequest, "아이디가 비어있습니다")
		return
	}

	// orgMoney := user.Money

	// err = db.DB.Model(&model.User{}).Select("money").Where("id = ?", user.Id).Find(user).Error
	// if err != nil {
	// 	log.Println(err)
	// 	module.Response(c, http.StatusBadRequest, "아이디가 비어있습니다")
	// 	return
	// }

	// user.Money += orgMoney

	err = db.DB.Model(&model.User{}).Where("id = ?", user.Id).Update("money", user.Money).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "save money update error")
		return
	}

	module.Response(c, http.StatusAccepted, "save money update success")
}
