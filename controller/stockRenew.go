package controller

import (
	"log"
	"net/http"

	"example.com/m/v2/module"
	"github.com/gin-gonic/gin"
)

func StockRenew(c *gin.Context) {

	err := module.Crawling()
	if err != nil {
		log.Println(err)
		module.Response(c, http.StatusBadRequest, "stock renew crawling error")
		return
	}

	module.Response(c, http.StatusAccepted, "stock renew success")
}
