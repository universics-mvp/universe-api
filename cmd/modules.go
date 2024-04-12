package cmd

import (
	route "main/internal/application"
	"main/internal/application/question_controller"
	"main/internal/config"
	challenge_infrastructure "main/internal/infrastructure/dailyChallenge"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	route.Module,
	question_controller.Module,
	challenge_infrastructure.Module,
)
