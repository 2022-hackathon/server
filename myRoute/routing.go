package myroute

import (
	"example.com/m/v2/controller"
	"example.com/m/v2/middleware"

	"github.com/gin-gonic/gin"
)

func Routing(c *gin.Engine) {

	c.Use(middleware.CORSMiddleware())

	c.POST("/signup", controller.SignUp)
	c.POST("/login", controller.Login)
	c.POST("/addboard", controller.AddBoard)
	c.POST("/deleteboard", controller.DelBoard)
	c.POST("/addcategorie", controller.AddCategorie)
	c.POST("/savemoney", controller.SaveMoney)

	c.GET("/recommandrandom", controller.RecommandRandom) // -----------------------
	c.GET("/recommandnubrank", controller.RecommandNubRank)
	c.GET("/getrank", controller.GetRank)
	c.GET("/stockrank", controller.GetStock)
	c.GET("/renew", controller.StockRenew)
	c.GET("/gamegetuser", controller.GameGetUser)
	c.GET("/getmycategorie", controller.GetMyCategorie)
	c.GET("/getmyboard", controller.GetMyBoard)
	c.GET("/", middleware.TokenAuthMiddleware())

}
