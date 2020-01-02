package routers

import (
	"github.com/gin-gonic/gin"
	"FirstGin/middleware/jwt"
	"FirstGin/pkg/setting"
	"FirstGin/routers/api"
	"FirstGin/routers/api/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//r.GET("/test", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "test",
	//	})
	//})
	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
