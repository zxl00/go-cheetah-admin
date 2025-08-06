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

func Menu(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysMenu().Create)
		r.POST("/delete/:id", apiSystem.NewSysMenu().Delete)
		r.POST("/update/:id", apiSystem.NewSysMenu().Update)
		r.GET("/list", apiSystem.NewSysMenu().List)
		r.GET("/get/:id", apiSystem.NewSysMenu().Get)
	}
	return r
}
