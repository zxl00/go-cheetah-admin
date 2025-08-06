/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package middles

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LogHandlerFunc() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义格式
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	})
}
