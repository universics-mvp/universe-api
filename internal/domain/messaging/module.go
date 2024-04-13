package messaging

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewChatService),
	fx.Provide(NewIMService),
)
