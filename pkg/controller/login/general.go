/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

/* UserGeneralLogin*/

package login

import (
	"errors"

	reqLogin "github.com/zxl00/go-cheetah-admin/internal/model/request/login"
	"github.com/zxl00/go-cheetah-admin/pkg/controller/system"
	"github.com/zxl00/go-cheetah-admin/pkg/global"
)

func (sl *sysLogin) GeneralLogin(request *reqLogin.ReqLogin) (error, interface{}) {
	ok, userInfo := system.NewSysUser(sl.ctx).GetByUsernameAndPwd(request.Username, request.Password)
	if ok && userInfo != nil {
		return nil, userInfo
	}
	return global.OtherErr(errors.New("登录失败")), nil
}
