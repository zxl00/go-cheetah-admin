/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import (
	"github.com/zxl00/go-cheetah-admin/internal/model"
)

// Ldap 用户登录ldap配置

type Ldap struct {
	model.BaseModel
	Address   string     `gorm:"column:address;type:varchar(256);not null" json:"address"`
	DN        string     `gorm:"column:dn" json:"dn"`
	AdminUser string     `gorm:"column:admin_user;not null" json:"admin_user"`
	Password  string     `gorm:"column:password" json:"password"`
	OU        string     `gorm:"column:ou" json:"ou"`
	Filter    string     `gorm:"column:filter;not null" json:"filter"`
	Mapping   model.JSON `gorm:"column:mapping;type:json;comment:'属性映射'" json:"mapping"`
	SSL       uint       `gorm:"type:tinyint(1);default:2;comment:'状态(正常/禁用, 默认禁用)'" json:"ssl"`
	Status    uint       `gorm:"type:tinyint(1);default:2;comment:'状态(正常/禁用, 默认禁用)'" json:"status"`
}

func (*Ldap) TableName() string {
	return "sys_ldap"
}
