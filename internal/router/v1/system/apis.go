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

func Apis(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create", apiSystem.NewSysApis().Create)
		r.POST("/delete/:id", apiSystem.NewSysApis().Delete)
		r.POST("/update/:id", apiSystem.NewSysApis().Update)
		r.GET("/list", apiSystem.NewSysApis().List)
		r.GET("/get/:id", apiSystem.NewSysApis().Get)
		r.GET("/get/group", apiSystem.NewSysApis().GetApiGroup)
	}
	return r
}
