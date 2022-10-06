package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/helper"
	"example.com/m/v2/module"

	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	user := &model.User{}

	err := c.Bind(user)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "login bind error")
		return
	}

	if user.Id == "" || user.Pw == "" {
		log.Println("아이디나 페스워드가 비어있습니다")
		module.Response(c, http.StatusBadRequest, "아이디나 페스워드가 비어있습니다")
		return
	}

	user.Pw, err = helper.HashPassword(user.Pw)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "login hash pwd error")
		return
	}

	err = db.DB.Model(&model.User{}).Where("id = ? and pw = ?", user.Id, user.Pw).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "아이디나 비밀번호를 확인해야함")
		return
	}

	tk, err := helper.CreateJWT(user.Id)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "login create token error")
		return
	}

	err = helper.StoreAuth(tk)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "login auth a error")
		return
	}

	module.ResponseToken(c, http.StatusAccepted, tk.AccessToken, user.Id, user.NickName)
}
