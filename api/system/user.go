/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"strconv"

	"github.com/gin-gonic/gin"

	reqSystem "github.com/zxl00/go-cheetah-admin/internal/model/request/system"
	"github.com/zxl00/go-cheetah-admin/pkg/controller/system"
	"github.com/zxl00/go-cheetah-admin/pkg/global"
	"github.com/zxl00/go-cheetah-admin/pkg/utils"
)

type UserInterface interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type sysUser struct{}

func NewSysUser() UserInterface {
	return &sysUser{}
}

func (su *sysUser) Create(ctx *gin.Context) {
	body := new(reqSystem.CreateUserReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	utils.TagAes(body)
	if err := system.NewSysUser(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建成功", nil)
}
func (su *sysUser) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := system.NewSysUser(ctx).Delete(id); err != nil {
		global.ReturnContext(ctx).Failed("删除失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("删除成功", nil)
}

func (su *sysUser) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := new(reqSystem.UpdateUserReq)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	utils.TagAes(body)
	if err := system.NewSysUser(ctx).Update(id, body); err != nil {
		global.ReturnContext(ctx).Failed("更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("更新成功", nil)
}

func (su *sysUser) List(ctx *gin.Context) {
	params := new(struct {
		UserName string `form:"username"`
		Limit    int    `form:"limit"`
		Page     int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err, data := system.NewSysUser(ctx).List(params.UserName, params.Limit, params.Page); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}

func (su *sysUser) Get(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err, data := system.NewSysUser(ctx).Get(id); err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	} else {
		global.ReturnContext(ctx).Successful("查询成功", data)
	}
}
