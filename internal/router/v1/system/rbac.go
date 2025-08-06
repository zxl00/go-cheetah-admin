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

func RBAC(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/create/:id", apiSystem.NewSysRBAC().Create)
		r.GET("/role/get/:id", apiSystem.NewSysRBAC().GetRbacByRoleID)
	}
	return r
}
