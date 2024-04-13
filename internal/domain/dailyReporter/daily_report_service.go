package daily_reporter

import (
	"bytes"
	"fmt"
	"main/internal/domain/group"
	language_model_domain "main/internal/domain/languageModel"
	"main/internal/domain/message"
	"main/internal/domain/messaging"
	"main/pkg"
	"time"
)

const DAILY = "18 0 0 * * *"

const BASE_DAILY_REPORT_PROMPT = `Ты - помощник куратора первокурсников в университете, ниже представлена беседа студентов за последний день, проанализируй написанное и составь отчёт о проблемах, на которые куратору стоило бы обратить внимание:`

type DailyReportService struct {
	logger    pkg.Logger
	cron      pkg.CronRunner
	groupRepo group.GroupRepository
	msgRepo   message.MessageRepository
	gpt       language_model_domain.LanguageModel
	tgbot     messaging.TGBot
}

func NewDailyReportService(logger pkg.Logger, cronRunner pkg.CronRunner, groupRepo group.GroupRepository, msgRepo message.MessageRepository, gpt language_model_domain.LanguageModel, tgbot messaging.TGBot) DailyReportService {
	return DailyReportService{
		logger:    logger,
		cron:      cronRunner,
		groupRepo: groupRepo,
		msgRepo:   msgRepo,
		gpt:       gpt,
		tgbot:     tgbot,
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
	messages, err := dr.msgRepo.GetMessagesForChat(group.ChatID, startTime)
	if err != nil {
		dr.logger.Error(err)
		return err
	}
	prompt := generatePrompt(messages)
	messageText, err := dr.gpt.GetAnswer(prompt, BASE_DAILY_REPORT_PROMPT, 0.3)
	if err != nil {
		dr.logger.Error(err)
		return err
	}
	err = dr.tgbot.SendMessage(group.CuratorID, fmt.Sprintf("Подготовил ежедневный отчёт о группе %s, вот, какие проблемы студентов я обнаружил", group.Title))
	if err != nil {
		dr.logger.Error(err)
		return err
	}
	err = dr.tgbot.SendMessage(group.CuratorID, messageText)
	if err != nil {
		dr.logger.Error(err)
		return err
	}
	return nil
}

func generatePrompt(messages []message.Message) string {
	var buffer bytes.Buffer
	for _, message := range messages {
		time := time.Unix(message.Date, 0)
		buffer.WriteString(fmt.Sprintf("%s %s написал '%s',", time.Format("[15:04]"), message.UserFullName, message.Text))
	}
	return buffer.String()
}
