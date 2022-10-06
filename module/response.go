package module

import (
	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, msg interface{}) {

	if v, ok := msg.(string); ok {
		ifString(c, code, v)
	} else if v, ok := msg.(*model.User); ok {
		ifUser(c, code, v)
	} else if v, ok := msg.([]model.Board); ok {
		ifBoardSlice(c, code, v)
	} else if v, ok := msg.([]model.Categorie); ok {
		ifCategorieSlice(c, code, v)
	} else if v, ok := msg.([]model.StockInfo); ok {
		ifStockSlice(c, code, v)
	}
}

func ResponseToken(c *gin.Context, code int, token string, id string, nickname string) {

	c.JSON(code, gin.H{
		"status":   code,
		"token":    token,
		"id":       id,
		"nickname": nickname,
	})
}

func ifString(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status":  code,
		"message": msg,
	})
}

func ifUser(c *gin.Context, code int, msg *model.User) {
	c.JSON(code, gin.H{
		"status": code,
		"user":   msg,
	})
}

func ifStockSlice(c *gin.Context, code int, boards []model.StockInfo) {
	c.JSON(code, gin.H{
		"status": code,
		"data":   boards,
	})
}

func ifBoardSlice(c *gin.Context, code int, boards []model.Board) {
	c.JSON(code, gin.H{
		"status": code,
		"data":   boards,
	})
}

func ifCategorieSlice(c *gin.Context, code int, boards []model.Categorie) {
	c.JSON(code, gin.H{
		"status": code,
		"data":   boards,
	})
}
