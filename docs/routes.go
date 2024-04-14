package docs

import (
	"main/pkg"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type SwaggerRoutes struct {
	logger  pkg.Logger
	handler pkg.RequestHandler
}

func NewSwaggerRoutes(logger pkg.Logger, handler pkg.RequestHandler) SwaggerRoutes {
	return SwaggerRoutes{
		logger:  logger,
		handler: handler,
	}
}

func (r SwaggerRoutes) Setup() {
	r.handler.Gin.GET("/api/doc/*any", func(context *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler)(context)
	})

	SwaggerInfo.Title = "Universe API"
}

var Module = fx.Options(
	fx.Provide(NewSwaggerRoutes),
)
