package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ibrat-muslim/blog_app_api_gateway/api/v1"
	"github.com/ibrat-muslim/blog_app_api_gateway/config"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	grpcPkg "github.com/ibrat-muslim/blog_app_api_gateway/pkg/grpc_client"

	_ "github.com/ibrat-muslim/blog_app_api_gateway/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg        *config.Config
	GrpcClient grpcPkg.GrpcClientI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @host      localhost:8000
// @BasePath  /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:        opt.Cfg,
		GrpcClient: opt.GrpcClient,
	})

	apiV1 := router.Group("/v1")

	apiV1.POST("/auth/register", handlerV1.Register)

	apiV1.GET("/users/:id", handlerV1.GetUser)
	// apiV1.GET("/users/me", handlerV1.GetUserProfile)
	apiV1.GET("/users", handlerV1.GetUsers)
	apiV1.GET("/users/email/:email", handlerV1.GetUserByEmail)
	apiV1.POST("/users", handlerV1.CreateUser)
	apiV1.PUT("/users/:id", handlerV1.UpdateUser)
	apiV1.DELETE("users/:id", handlerV1.DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
