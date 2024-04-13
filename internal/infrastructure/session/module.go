package session_infrastructure

import (
	"main/internal/domain/session"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewSessionRepository),
	fx.Provide(session.NewSessionService),
)
