package daily_reporter

import (
	"main/internal/domain/group"
	"main/internal/domain/message"
	"main/pkg"
	"time"
)

const DAILY = "@daily"

type DailyReportService struct {
	logger      pkg.Logger
	cron        pkg.CronRunner
	groupRepo   group.GroupRepository
	messageRepo message.MessageRepository
}

func NewDailyReportService(logger pkg.Logger, cronRunner pkg.CronRunner, repo group.GroupRepository) DailyReportService {
	return DailyReportService{
		logger:    logger,
		cron:      cronRunner,
		groupRepo: repo,
	}
}

func (dr DailyReportService) Run() {
	dr.cron.AddFunc(DAILY, dr.sendDailyReports)
	dr.cron.Run()
}

func (dr DailyReportService) sendDailyReports() {
	groups, err := dr.groupRepo.List()
	if err != nil {
		dr.logger.Error(err)
		return
	}
	for _, group := range groups {
		dr.sendDailyReport(group)
	}
}

func (dr DailyReportService) sendDailyReport(group group.Group) error {
	startTime := time.Now().AddDate(0, 0, -1).Unix()
	messages, err := dr.messageRepo.GetMessagesForChat(group.ChatID, startTime)
	if err != nil {
		dr.logger.Error(err)
		return err
	}

}
