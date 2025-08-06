/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package logins

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	apiLogin "github.com/zxl00/go-cheetah-admin/api/logins"
)

func Login(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/ldap", authMiddleware.LoginHandler)
		r.POST("/general", authMiddleware.LoginHandler)
	}
	return r
}

func Resource(r *gin.RouterGroup) gin.IRoutes {
	{
		r.GET("/info", apiLogin.NewSysLogin().GetLoginUserResource)
	}
	return r
}
