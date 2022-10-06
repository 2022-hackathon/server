package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func RecommandRandom(c *gin.Context) {

	stocks := []model.StockInfo{}

	err := db.DB.Raw("select * from stock_infos order by rand()").Find(&stocks).Error
	if err != nil {
		log.Println("recommand find random")
		module.Response(c, http.StatusBadRequest, "recommand random find error")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
		"data":   stocks,
	})
}
