/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/zxl00/go-cheetah-admin/pkg/controller/system"
	"github.com/zxl00/go-cheetah-admin/pkg/global"
)

type SysRBAC interface {
	Create(ctx *gin.Context)
	GetRbacByRoleID(ctx *gin.Context)
}
type sysRbac struct {
}

func NewSysRBAC() SysRBAC {
	return &sysRbac{}
}

func (sr *sysRbac) Create(ctx *gin.Context) {
	body := new(struct {
		ApisID []int `json:"apis_id"`
	})
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysRBAC(ctx).Create(body.ApisID, roleID); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}

func (sr *sysRbac) GetRbacByRoleID(ctx *gin.Context) {
	roleID, _ := strconv.Atoi(ctx.Param("id"))
	roleIDs := system.NewSysRBAC(ctx).GetRbacByRoleID(roleID)
	global.ReturnContext(ctx).Successful("获取已经授权角色成功", roleIDs)
}
