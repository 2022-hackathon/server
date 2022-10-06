package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func AddBoard(c *gin.Context) {

	board := &model.Board{}

	err := c.Bind(board)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "add board bind error")
		return
	}

	if board.Title == "" || board.UserId == "" || board.BoardCategorie == "" {
		module.Response(c, http.StatusBadRequest, "제목이나 유저아이디가 비어있습니다")
		return
	}

	res := db.DB.Model(&model.User{}).Where("id = ?", board.UserId).Find(&model.User{})
	if res.RowsAffected != 1 {
		log.Println("add board id rowsAffected : ", res.RowsAffected)
		module.Response(c, http.StatusBadRequest, "아이디가 존재하지 않습니다")
		return
	}

	res = db.DB.Model(&model.Categorie{}).Where("User_categorie = ?", board.BoardCategorie).Find(&model.Categorie{})
	if res.RowsAffected != 1 {
		log.Println("add board cate rowsAffected : ", res.RowsAffected)
		module.Response(c, http.StatusBadRequest, "키테고리가 존재하지 않습니다")
		return
	}

	err = db.DB.Model(&model.Board{}).Create(board).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "add board create error")
		return
	}

	module.Response(c, http.StatusAccepted, "add board success")
}
