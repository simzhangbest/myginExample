package routers

import (
	"github.com/gin-gonic/gin"
	"myginExample/middleware/jwt"
	"myginExample/pkg/setting"
	"myginExample/pkg/upload"
	"myginExample/routers/api"
	"myginExample/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTags)

		apiv1.PUT("/tags/:id", v1.EditTag)

		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)

		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return r
}
