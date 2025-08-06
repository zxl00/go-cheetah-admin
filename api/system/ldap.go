/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"github.com/gin-gonic/gin"

	modSys "github.com/zxl00/go-cheetah-admin/internal/model/system"
	"github.com/zxl00/go-cheetah-admin/pkg/controller/system"
	"github.com/zxl00/go-cheetah-admin/pkg/global"
)

type LdapInterface interface {
	Create(ctx *gin.Context)
	Info(ctx *gin.Context)
	Ping(ctx *gin.Context)
}

type sysLdap struct{}

func NewSysLdap() LdapInterface {
	return &sysLdap{}
}

func (s *sysLdap) Create(ctx *gin.Context) {
	body := new(modSys.Ldap)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	if err := system.NewSysLdap(ctx).Create(body); err != nil {
		global.ReturnContext(ctx).Failed("创建或更新失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("创建或更新成功", nil)
}

func (s *sysLdap) Info(ctx *gin.Context) {
	err, ldap := system.NewSysLdap(ctx).Info()
	if err != nil {
		global.ReturnContext(ctx).Failed("查询失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("查询成功", ldap)
}

func (s *sysLdap) Ping(ctx *gin.Context) {
	body := new(modSys.Ldap)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.ReturnContext(ctx).Failed("参数错误", err.Error())
		return
	}
	err := system.NewSysLdap(ctx).Ping(body)
	if err != nil {
		global.ReturnContext(ctx).Failed("连接失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("连接成功", nil)
}
