package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func AddCategorie(c *gin.Context) {

	categorie := &model.Categorie{}

	err := c.Bind(categorie)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "add categorie bind error")
		return
	}

	if categorie.UserId == "" || categorie.UserCategorie == "" {
		log.Println("유저 아이디나 카테고리가 비어있음")
		module.Response(c, http.StatusBadRequest, "유저 아이디나 카테고리가 비어있음")
		return
	}

	err = db.DB.Model(&model.Categorie{}).Create(categorie).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "add categorie create error")
		return
	}

	module.Response(c, http.StatusAccepted, "add categorie success")
}
