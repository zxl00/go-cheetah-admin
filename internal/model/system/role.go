/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import "github.com/zxl00/go-cheetah-admin/internal/model"

type Role struct {
	model.BaseModel
	Name   string `gorm:"column:name;comment:'角色名称';size:128" json:"name"`
	Desc   string `gorm:"column:desc;comment:'角色描述';size:128" json:"desc"`
	Status uint   `gorm:"type:tinyint(1);default:1;comment:'用户状态(正常/禁用, 默认正常)'" json:"status"`
	Users  []User `gorm:"many2many:sys_user_role;" json:"users"  copier:"-"`
	Menus  []Menu `gorm:"many2many:sys_role_menu;" json:"menus"  copier:"-"`
}

func (*Role) TableName() string {
	return "sys_role"
}
