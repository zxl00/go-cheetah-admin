/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package response

type CasbinPolicy struct {
	PType  string `gorm:"column:ptype" json:"p_type"`
	RoleID string `gorm:"column:v0" json:"V0"`
	Path   string `gorm:"column:v1" json:"V1"`
	Method string `gorm:"column:v2" json:"V2"`
	Desc   string `gorm:"column:v3" json:"V3"`
}
