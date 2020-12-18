package router

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go_gateway_demo/controller"
	"go_gateway_demo/docs"
	"go_gateway_demo/middleware"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// admin_login的路由注册，将group注册到router，管理员登陆，用redis保存cookie
	adminLoginRouter := router.Group("/admin_login")
	// 设置router中间件，分布式持久化-redis
	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalf("sessions.NewRedisStore err:%v", err)
	}
	adminLoginRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		// 子group的注册
		controller.AdminLoginRegister(adminLoginRouter)
	}

	// admin_info的路由注册，将group注册到router
	adminRouter := router.Group("/admin")
	adminRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		// 子路由的注册
		controller.AdminRegister(adminRouter)
	}

	// service的路由注册
	serviceRouter := router.Group("/service")
	serviceRouter.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		// 子路由的注册
		controller.ServiceRegister(serviceRouter)
	}

	return router
}
