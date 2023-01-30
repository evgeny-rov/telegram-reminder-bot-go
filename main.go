package main

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

type Reminder struct {
	chatId   int64
	duration time.Duration
	firesAt  time.Time
	timer    *time.Timer
}

var db = make(map[int64]Reminder)

func response(lang string) translation {
	switch lang {
	case "en":
		return Translations.en
	case "ru":
		return Translations.ru
	default:
		return Translations.en
	}
}

func handleStart(c tele.Context) error {
	return c.Send(response(c.Sender().LanguageCode).start)
}

func handleHelp(c tele.Context) error {
	return c.Send(response(c.Sender().LanguageCode).help)
}

func handleNewReminder(c tele.Context) error {
	lang := c.Sender().LanguageCode
	timeRegExp := regexp.MustCompile(`^\d?\d?-?\d?\d`)
	timeArg := strings.TrimSpace(timeRegExp.FindString(c.Message().Payload))

	if len(timeArg) == 0 {
		return c.Send(response(lang).badParams)
	}

	timeWithUnits := strings.ReplaceAll(timeArg, "-", "h") + "m"
	duration, err := time.ParseDuration(timeWithUnits)
	firesAt := time.Now().Add(duration)

	if err != nil {
		return c.Send(response(lang).badParams)
	}

	const minutesInADay = 1440

	if duration.Minutes() < 1 || duration.Minutes() > float64(minutesInADay) {
		return c.Send(response(lang).outOfRange)
	}

	chatId := c.Chat().ID

	if prevReminder, hasPrevReminder := db[chatId]; hasPrevReminder {
		prevReminder.timer.Stop()
		delete(db, chatId)
	}

	timer := time.AfterFunc(duration, func() {
		messageArg := c.Message().Payload[len(timeArg):]
		message := strings.TrimSpace(messageArg)

		if len(message) > 0 {
			c.Send(response(lang).alertWithMessage + " " + message)
		} else {
			c.Send(response(lang).alertWithoutMessage)
		}

		delete(db, chatId)
	})

	db[chatId] = Reminder{duration: duration, firesAt: firesAt, chatId: chatId, timer: timer}
	return c.Send(response(lang).created)
}

func handleCancelReminder(c tele.Context) error {
	chatId := c.Chat().ID

	if reminder, hasReminder := db[chatId]; hasReminder {
		reminder.timer.Stop()
		delete(db, chatId)

		return c.Send(response(c.Sender().LanguageCode).cancelled)
	}

	return c.Send(response(c.Sender().LanguageCode).noReminders)
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token, ok := os.LookupEnv("TELEGRAM_BOT_API_TOKEN")

	if !ok {
		log.Fatal("token not set")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 60 * time.Second},
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", handleStart)
	bot.Handle("/help", handleHelp)
	bot.Handle("/remind", handleNewReminder)
	bot.Handle("/cancel", handleCancelReminder)

	bot.Start()
}
