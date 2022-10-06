package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/db"
	"example.com/m/v2/model"
	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func GetStock(c *gin.Context) {

	stocks := []model.StockInfo{}

	/*err := module.Crawling()
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get stock crawling error")
		return
	}*/

	err := db.DB.Model(&model.StockInfo{}).Omit("per").Find(&stocks).Error
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "get stock find error")
		return
	}

	module.Response(c, http.StatusAccepted, stocks)
}
