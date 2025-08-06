/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package login

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"  aes:"true" binding:"required"`
}
