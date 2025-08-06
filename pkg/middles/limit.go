/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package middles

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"

	"github.com/zxl00/go-cheetah-admin/pkg/global"
)

func RateLimitMiddle(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(ctx *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			global.ReturnContext(ctx).Failed("failed", "访问限流")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
