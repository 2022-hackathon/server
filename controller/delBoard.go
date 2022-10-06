package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func DelBoard(c *gin.Context) {

	board := &model.Board{}

	err := c.Bind(board)
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "delete board ")
		return
	}

	err = db.DB.Model(&model.Board{}).Where("user_id = ? and title = ?", board.UserId, board.Title).Delete(board).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "delete board delete error")
		return
	}

	module.Response(c, http.StatusAccepted, "delete board success")
}
