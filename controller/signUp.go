package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/helper"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	user := &model.User{}

	err := c.Bind(user)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "signup bind error")
		return
	}

	if user.Id == "" || user.Pw == "" {
		module.Response(c, http.StatusBadRequest, "아이디나 비밀번호가 비어있음")
		return
	}

	db_res := db.DB.Model(&model.User{}).Where("id = ?", user.Id).Find(&model.User{})
	if db_res.RowsAffected != 0 {
		log.Println("아이디가 이미 존재합니다")
		module.Response(c, http.StatusBadRequest, "아이디가 이미 존재합니다")
		return
	}

	user.Pw, err = helper.HashPassword(user.Pw)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "signup hash pwd error")
		return
	}

	err = db.DB.Model(&model.User{}).Create(user).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "signup create error")
		return
	}

	module.Response(c, http.StatusAccepted, "signp success")
}
