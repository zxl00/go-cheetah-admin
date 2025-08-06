/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

type CreateAPIsReq struct {
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	Desc     string `json:"desc"  binding:"required"`
	ApiGroup string `json:"apiGroup" binding:"required"`
	//Menus    []int  `json:"menus"`
}

type UpdateAPIsReq struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Desc     string `json:"desc"`
	ApiGroup string `json:"apiGroup"`
	//Menus    []int  `json:"menus"`
}
