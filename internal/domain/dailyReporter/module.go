package daily_reporter

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDailyReportService),
)
