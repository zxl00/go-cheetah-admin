/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package login

import (
	"context"

	reqLogin "github.com/zxl00/go-cheetah-admin/internal/model/request/login"
)

type SysLogin interface {
	LdapLogin(request *reqLogin.ReqLogin) (error, interface{})
	GeneralLogin(request *reqLogin.ReqLogin) (error, interface{})
}
type sysLogin struct {
	ctx context.Context
}

func NewSysLogin(ctx context.Context) SysLogin {
	return &sysLogin{ctx: ctx}
}
