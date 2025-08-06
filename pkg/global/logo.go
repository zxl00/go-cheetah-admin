/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package global

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitSysTips() {
	var logo = `
 __ __   ____   _____ ______  
|  |  \_/ ___\ /     \\____ \ 
|  |  /\  \___|  Y Y  \  |_> >
|____/  \___  >__|_|  /   __/ 
            \/      \/|__|
`
	var sys = `
Version: 1.0.0
Author: AnRuo`
	var run = "Address: " + fmt.Sprintf("%s:%d", viper.GetString("server.address"),
		viper.GetInt("server.port"))
	fmt.Println(logo)
	fmt.Println(sys)
	fmt.Println(run)

}
