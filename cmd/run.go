/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"github.com/zxl00/go-cheetah-admin/internal/router"
	"github.com/zxl00/go-cheetah-admin/pkg/global"
)

func Run() {
	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("server.address"),
			viper.GetInt("server.port")),
		Handler:        router.RegisterRouters(),
		MaxHeaderBytes: 1 << 20,
	}
	// 打印服务启动参数
	// 关闭服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 定时清理文件
	quit := make(chan os.Signal)
	// 获取停止服务信号，kill  -9 获取不到
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 执行延迟停止
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

}

func init() {
	global.InitConfig()
	global.InitSysTips()
	global.InitLog()
	global.InitMysql()
	global.InitCasbinEnforcer()
}
