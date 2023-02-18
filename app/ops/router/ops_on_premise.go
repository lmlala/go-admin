package router

import (
	"go-admin/app/ops/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	// routerCheckRole = append(routerCheckRole, registerSysUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, addOnPremiseRouter)
}

// 无需认证的路由代码
// func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
func addOnPremiseRouter(v1 *gin.RouterGroup) {
	api := apis.OpsOnPremise{}
	r := v1.Group("/ops/onpremise")
	{
		r.GET("", api.GetPage)
		// r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}
