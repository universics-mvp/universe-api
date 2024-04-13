package cmd

import (
	route "main/internal/application"
	"main/internal/application/categorization"
	challenge_answer_application "main/internal/application/challengeAnswer"
	challenge_application "main/internal/application/dailyChallenge"
	question_controller "main/internal/application/question"
	"main/internal/config"
	challenge_answer_domain "main/internal/domain/challengeAnswer"
	challenge_domain "main/internal/domain/dailyChallenge"
	challenge_answer_infrastructure "main/internal/infrastructure/challengeAnswer"
	challenge_infrastructure "main/internal/infrastructure/dailyChallenge"
	yandex_language_model "main/internal/infrastructure/yandexLanguageModel"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	route.Module,
	question_controller.Module,

	challenge_infrastructure.Module,
	challenge_domain.Module,
	challenge_application.Module,

	yandex_language_model.Module,

	challenge_answer_domain.Module,
	challenge_answer_infrastructure.Module,
	challenge_answer_application.Module,

	categorization.Module,
)
