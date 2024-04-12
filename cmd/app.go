package cmd

import (
	"context"
	route "main/internal/application"
	"main/internal/config"
	"main/pkg"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)


func Run() any {
	return func(
		route route.Routes,
		handler pkg.RequestHandler,
		env config.Env,
		logger pkg.Logger,
	) {
		route.Setup()
		
		err := handler.Gin.Run(":" + env.Port)
		if err != nil {
			logger.Error(err)
			return 
		}
	}
}


func StartApp() error {
	logger := pkg.GetLogger(config.NewEnv())
	opts := fx.Options(
		fx.WithLogger(func () fxevent.Logger  {
			return logger.GetFxLogger()
		}),
		fx.Invoke(Run()),
	)
	ctx := context.Background()
	app := fx.New(CommonModules, opts)
	err := app.Start(ctx)
	return err
}
