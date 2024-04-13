package cmd

import (
	"context"
	route "main/internal/application"
	chatbot "main/internal/application/chat_bot"
	"main/internal/config"
	daily_reporter "main/internal/domain/dailyReporter"
	"main/pkg"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Run() any {
	return func(
		route route.Routes,
		handler pkg.RequestHandler,
		tgHandler chatbot.TgHandler,
		dailyReporter daily_reporter.DailyReportService,
		env config.Env,
		logger pkg.Logger,
	) {
		route.Setup()
		tgHandler.Run()
		dailyReporter.Run()
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
		fx.WithLogger(func() fxevent.Logger {
			return logger.GetFxLogger()
		}),
		fx.Invoke(Run()),
	)
	ctx := context.Background()
	app := fx.New(CommonModules, opts)
	err := app.Start(ctx)
	return err
}
