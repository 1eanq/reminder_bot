package app

import (
	"fmt"
	"github.com/mymmrac/telego"
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

	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Button A").
					WithRequestPoll((tu.PollTypeRegular())),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("Contact").WithRequestContact(),
				tu.KeyboardButton("Location").WithRequestLocation(),
			),
		).WithResizeKeyboard().WithInputFieldPlaceholder("Select!!!")

		msg := tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Hello world",
		).WithReplyMarkup(keyboard).WithProtectContent()

		bot.SendMessage(msg)
	}

}
