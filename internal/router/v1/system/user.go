/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	apiSystem "github.com/zxl00/go-cheetah-admin/api/system"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysUser().Create)
		r.POST("/delete/:id", apiSystem.NewSysUser().Delete)
		r.POST("/update/:id", apiSystem.NewSysUser().Update)
		r.GET("/list", apiSystem.NewSysUser().List)
		r.GET("/get/:id", apiSystem.NewSysUser().Get)
		r.POST("/logout", authMiddleware.LogoutHandler)
		r.POST("/refresh", authMiddleware.RefreshHandler)
	}
	return r
}
