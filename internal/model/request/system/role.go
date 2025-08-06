/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package system

type CreateRoleReq struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status uint   `json:"status"`
	Users  []int  `json:"users"`
	Menus  []int  `json:"menus"`
}
