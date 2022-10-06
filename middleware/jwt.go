package middleware

import (
	"net/http"

	"example.com/m/v2/helper"
	"example.com/m/v2/module"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc { // token확인 middleware

	return func(c *gin.Context) {

		if err := helper.TokenVaild(c); err != nil {
			module.Response(c, http.StatusUnauthorized, "vaild error")
		}

		ad, err := helper.ExtractTokenMetadata(c)
		if err != nil {
			module.Response(c, http.StatusUnauthorized, "extract token error")
		}

		if _, err := helper.CheckAuth(ad); err != nil {
			module.Response(c, http.StatusUnauthorized, "fetch auth error")
		}

		c.Next()
	}
}
