package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func GetMyBoard(c *gin.Context) {

	boards := []model.Board{}

	id := c.Query("id")
	if id == "" {
		module.Response(c, http.StatusBadRequest, "get my board id가 비어있음")
		return
	}

	categorie := c.Query("categorie")
	if categorie == "" {
		module.Response(c, http.StatusBadRequest, "get my board cate가 비어있음")
		return
	}

	err := db.DB.Model(&model.Board{}).Where("user_id = ? and board_categorie = ?", id, categorie).Find(&boards).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get my board find error")
		return
	}
	module.Response(c, http.StatusAccepted, boards)
}
