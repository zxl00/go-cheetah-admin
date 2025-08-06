/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

type CreateMenuReq struct {
	Name      string `json:"name" binding:"required"`
	NameCode  string `json:"name_code"`
	IsShow    uint   `json:"is_show"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Sort      int    `json:"sort"`
	ParentId  uint   `json:"parent_id"`
	Component string `json:"component"`
	MenuType  uint   `json:"menu_type"`
	//APIs      []int  `json:"apis"`
}
