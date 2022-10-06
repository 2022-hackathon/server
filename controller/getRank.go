package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

type rankInfo struct {
	NickName string `json:"nickname"`
	Money    int    `json:"money"`
}

func GetRank(c *gin.Context) {

	rankInfos := []rankInfo{}

	err := db.DB.Model(&model.User{}).Select("nick_name", "money").Order("money DESC").Find(&rankInfos).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get rank find error")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
		"data":   rankInfos,
	})
}
