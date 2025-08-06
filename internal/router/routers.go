/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

// 路由

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zxl00/go-cheetah-admin/api"
	routerLogin "github.com/zxl00/go-cheetah-admin/internal/router/v1/logins"
	routerSys "github.com/zxl00/go-cheetah-admin/internal/router/v1/system"
	"github.com/zxl00/go-cheetah-admin/pkg/middles"
	"io"
	"net/http"
)

func RegisterRouters() *gin.Engine {
	r := gin.New()
	// 新增：动态环境变量端点（必须在静态文件中间件之前注册！）
	r.Use(middles.LogHandlerFunc(), gin.Recovery(), middles.Cors())
	r.Use(middles.Serve("/", middles.EmbedFolder(api.FS(), "ui/dist")))
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		panic(err)
	}
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		return
	})
	PrivateGroup := r.Group("")
	PrivateGroup.Use(authMiddleware.MiddlewareFunc(), middles.RbacMiddle())
	//PrivateGroup.Use()
	{
		UserGroup := PrivateGroup.Group("/sys/user")
		{
			routerSys.User(UserGroup, authMiddleware)
		}
		MenuGroup := PrivateGroup.Group("/sys/menu")
		{
			routerSys.Menu(MenuGroup)
		}
		RoleGroup := PrivateGroup.Group("/sys/role")
		{
			routerSys.Role(RoleGroup)
		}
		ApisGroup := PrivateGroup.Group("/sys/apis")
		{
			routerSys.Apis(ApisGroup)
		}
		LdapGroup := PrivateGroup.Group("/sys/ldap")
		{
			routerSys.Ldap(LdapGroup)
		}
		RbacGroup := PrivateGroup.Group("/sys/rbac")
		{
			routerSys.RBAC(RbacGroup)
		}
		LoginUserResource := PrivateGroup.Group("/sys/login")
		{
			routerLogin.Resource(LoginUserResource)
		}

	}

	PublicGroup := r.Group("")
	{
		LoginGroup := PublicGroup.Group("/sys/login")
		{
			routerLogin.Login(LoginGroup, authMiddleware)
		}
	}
	r.NoRoute(func(ctx *gin.Context) {
		fileObj, err := api.FS().Open("ui/dist/index.html")
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		data, err := io.ReadAll(fileObj)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})
	return r
}
