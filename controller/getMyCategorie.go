package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func GetMyCategorie(c *gin.Context) {

	id := c.Query("id")

	categorie := []model.Categorie{}

	if id == "" {
		log.Println("id가 비어있습니다")
		module.Response(c, http.StatusBadRequest, "id가 비어있습니다")
		return
	}

	err := db.DB.Model(&model.Categorie{}).Where("user_id = ?", id).Find(&categorie).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get my categorie find error")
		return
	}

	module.Response(c, http.StatusAccepted, categorie)
}
