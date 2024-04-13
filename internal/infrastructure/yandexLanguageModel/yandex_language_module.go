package yandex_language_model

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLanguageModel),
)
