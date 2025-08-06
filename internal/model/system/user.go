/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"github.com/zxl00/go-cheetah-admin/internal/model"
)

type User struct {
	model.BaseModel
	Username string `json:"userName" gorm:"index;not null;unique;comment:用户登录名"` // 用户登录名
	Password string `json:"password"  gorm:"comment:用户登录密码"`                     // 用户登录密码
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`           // 用户昵称
	Avatar   string `gorm:"column:avatar;comment:'用户头像';type:longtext" json:"avatar"`
	Email    string `gorm:"column:email;comment:'邮箱';size:128" json:"email"`
	Phone    string `gorm:"column:phone;comment:'手机号码';size:11" json:"phone"`
	Status   uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Roles    []Role `gorm:"many2many:sys_user_role;" json:"roles" copier:"-"`
}

func (*User) TableName() string {
	return "sys_user"
}
