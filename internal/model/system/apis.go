/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

import "github.com/zxl00/go-cheetah-admin/internal/model"

type APIs struct {
	model.BaseModel
	Path     string `gorm:"not null;unique" json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	Desc     string `json:"desc"  binding:"required"`
	ApiGroup string `json:"apiGroup" binding:"required"`
	//Menus    []Menu `gorm:"many2many:sys_menu_api;" json:"menus"`
}

func (*APIs) TableName() string {
	return "system_apis"
}
