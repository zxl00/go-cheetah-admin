/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package utils

import (
	"strconv"
	"time"
)

// 时间格式转换

func FormatNanSecond(nanoSecond string) string {
	atoi, err := strconv.Atoi(nanoSecond)
	if err != nil {
		return nanoSecond
	}
	microSeconds := int64(atoi) / 1e3
	return time.UnixMicro(microSeconds).Format("2006/01/02 15:04:05.999999")
}
