package pkg

import "github.com/robfig/cron"

type CronRunner struct {
	logger     Logger
	cronRunner *cron.Cron
}

func NewCronRunner(logger Logger) CronRunner {
	return CronRunner{logger: logger, cronRunner: cron.New()}
}

func (cr CronRunner) Run() {
	cr.cronRunner.Start()
}

func (cr CronRunner) AddFunc(timeTemplate string, function func()) {
	cr.cronRunner.AddFunc(timeTemplate, function)
}
