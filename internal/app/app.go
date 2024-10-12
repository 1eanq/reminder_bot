package app

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log/slog"
)

type Application struct {
	logger   *slog.Logger
	APIToken string
	dbPath   string
}

func New(
	log *slog.Logger,
	APIToken string,
	dbPath string,
) *Application {
	return &Application{log, APIToken, dbPath}
}

func Run(app *Application) {
	log := app.logger
	APIToken := app.APIToken
	dbPath := app.dbPath

	log.Info("Application started!")
	log.Info("Database path", slog.String("databasePath", dbPath))
	log.Info("API token", slog.String("APIToken", APIToken))

	bot, err := telego.NewBot(APIToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Error("Failed to create bot")
		panic(err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()

	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Send start message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(message.Chat.ID),
			fmt.Sprintf("Hello %s! Press REGISTER to start", message.From.FirstName),
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("REGISTER").WithCallbackData("register"),
			)),
		))
	}, th.CommandEqual("start"))

	// Register new handler with match on the call back query
	// with data equal to `go` and non-nil message
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "In order to register you need to send your phone number").WithReplyMarkup(
			tu.Keyboard(tu.KeyboardRow(tu.KeyboardButton("Send number").WithRequestContact())).
				WithResizeKeyboard().WithInputFieldPlaceholder("Send your phone number"),
		))

		// Answer callback query
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("register"))

	bh.Start()
}
