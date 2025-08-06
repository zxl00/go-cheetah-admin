/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"github.com/gin-gonic/gin"

	apiSystem "github.com/zxl00/go-cheetah-admin/api/system"
)

func Role(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysRole().Create)
		r.POST("/delete/:id", apiSystem.NewSysRole().Delete)
		r.POST("/update/:id", apiSystem.NewSysRole().Update)
		r.GET("/list", apiSystem.NewSysRole().List)
		r.GET("/get/:id", apiSystem.NewSysRole().Get)
	}
	return r
}
