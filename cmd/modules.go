package cmd

import (
	route "main/internal/application"
	"main/internal/application/categorization"
	challenge_answer_application "main/internal/application/challengeAnswer"
	chatbot "main/internal/application/chat_bot"
	challenge_application "main/internal/application/dailyChallenge"
	question_controller "main/internal/application/question"
	"main/internal/config"
	challenge_answer_domain "main/internal/domain/challengeAnswer"
	challenge_domain "main/internal/domain/dailyChallenge"
	"main/internal/domain/messaging"
	challenge_answer_infrastructure "main/internal/infrastructure/challengeAnswer"
	challenge_infrastructure "main/internal/infrastructure/dailyChallenge"
	session_infrastructure "main/internal/infrastructure/session"
	tgbot "main/internal/infrastructure/tgBot"
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

	messaging.Module,
	session_infrastructure.Module,
	tgbot.Module,
	chatbot.Module,
	categorization.Module,
)
