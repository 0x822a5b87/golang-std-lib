package jwt

import (
	setting "gin/blog/pkg"
	"gin/blog/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = setting.SUCCESS
		token := c.Query("token")
		claims, err := util.ParseToken(token)
		if err != nil {
			code = setting.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = setting.ErrorAuthCheckTokenTimeout
		}

		if code != setting.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  setting.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
