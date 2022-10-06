package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func GameGetUser(c *gin.Context) {

	user := &model.User{}

	id := c.Query("id")

	if id == "" {
		log.Println("id가 비어있습니다")
		module.Response(c, http.StatusBadRequest, "id가 비어있습니다")
		return
	}

	err := db.DB.Model(&model.User{}).Omit("pw").Where("id = ?", id).Find(user).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "game get user find error")
		return
	}

	module.Response(c, http.StatusAccepted, user)
}
