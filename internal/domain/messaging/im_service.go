package messaging

import (
	"fmt"
	"log"
	"main/internal/application/categorization"
	question_controller "main/internal/application/question"
	language_model_domain "main/internal/domain/languageModel"
	"main/internal/domain/message"
	"main/internal/domain/session"
	"main/pkg"
	"os"
	"path/filepath"

	"gopkg.in/telebot.v3"
)

const BASE_IM_PROMPT = "Ты - бот для помощи студентам первокурсникам в адаптации, вот необходимая информация: "
const SCHEDULE_MSG_TEXT = "Если я правильно понял, тебя интересует расписание, посмотреть его ты можешь здесь: https://guap.ru/rasp/"

type IMService struct {
	bot         TGBot
	logger      pkg.Logger
	msgRepo     message.MessageRepository
	sessService session.SessionService
	gpt         language_model_domain.LanguageModel
	categorizer categorization.Categorizer
}

func NewIMService(logger pkg.Logger, bot TGBot, msgRepo message.MessageRepository, sessService session.SessionService, gpt language_model_domain.LanguageModel, categorizer categorization.Categorizer) IMService {
	return IMService{
		bot:         bot,
		logger:      logger,
		sessService: sessService,
		msgRepo:     msgRepo,
		gpt:         gpt,
		categorizer: categorizer,
	}
}

func (svc IMService) HandleMessage(receivedMessage telebot.Message) error {
	chatSession, err := svc.sessService.GetOrCreateSessionForChat(receivedMessage.Chat.ID)
	if err != nil {
		return err
	}
	message := message.FromTGMessage(receivedMessage, *chatSession.ID)
	_, err = svc.msgRepo.Save(&message)
	if err != nil {
		return err
	}

	categories, err := svc.categorizer.Categorize(receivedMessage.Text, question_controller.StandardCategoies)
	if err != nil {
		return err
	}
	if includes(categories, "Расписание") {
		return svc.bot.SendMessage(receivedMessage.Chat.ID, SCHEDULE_MSG_TEXT)
	}

	knowledgeBase, err := getKnoledgeBase()
	if err != nil {
		return err
	}

	prompt := BASE_IM_PROMPT + knowledgeBase

	answer, err := svc.gpt.GetAnswer(receivedMessage.Text, prompt, 0.3)

	if err != nil {
		return err
	}

	return svc.bot.SendMessage(receivedMessage.Chat.ID, answer)
}

func getKnoledgeBase() (string, error) {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Fatal(err)
	}
	knowledge_base, err := os.ReadFile(filepath.Join(dir, "knowledge_base.txt"))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	return string(knowledge_base), nil
}

func includes(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
