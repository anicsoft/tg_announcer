package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

const (
	_initDataKey contextKey = "init-data"
)

const endpoint = "/notify"

type contextKey string

func main() {
	log.Println("Starting API service")

	webAppURL := os.Getenv("TELEGRAM_WEB_APP_URL")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := gotgbot.NewBot(botToken, nil)
	if err != nil {
		log.Fatalf("Telegram Bot API initialization error: %v", err)
	}
	log.Println("Telegram Bot API initialized")

	r := gin.New()

	r.Use(gin.Logger())
	r.POST("/bot", CreateBotEndpointHandler(bot, webAppURL))
	r.Use(authMiddleware(botToken))
	r.GET("/", showInitDataMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(r.Run(":" + port))
}

func showInitDataMiddleware(context *gin.Context) {
	initData, ok := ctxInitData(context.Request.Context())
	if !ok {
		context.AbortWithStatusJSON(401, map[string]any{
			"message": "Init data not found",
		})
		return
	}

	context.JSON(http.StatusOK, initData)
}

// According to the https://core.telegram.org/bots/api#setwebhook webhook will receive JSON-serialized Update structure
// Handler created by this function parses Update structure and replies to any message with welcome text and inline keyboard to open Mini App
func CreateBotEndpointHandler(bot *gotgbot.Bot, appURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		update := gotgbot.Update{}
		err := c.Bind(&update)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if update.Message == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("bot update didn't include a message")})
			return
		}

		switch update.Message.Text {
		case "/start":
			userData := &gin.H{
				"id":            update.Message.From.Id,
				"first_name":    update.Message.From.FirstName,
				"last_name":     update.Message.From.LastName,
				"username":      update.Message.From.Username,
				"language_code": update.Message.From.LanguageCode,
			}

			go func() {
				if err := sendUserData(userData); err != nil {
					log.Println("failed to send user data:", err)
				}
			}()
			launchWebApp(appURL, c, bot, update)
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func launchWebApp(appURL string, c *gin.Context, bot *gotgbot.Bot, update gotgbot.Update) {
	message := "Launch application"
	opts := &gotgbot.SendMessageOpts{
		ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					gotgbot.InlineKeyboardButton{
						Text: "Open WebApp", WebApp: &gotgbot.WebAppInfo{Url: appURL},
					},
				},
			},
		},
	}

	if _, err := bot.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func sendUserData(data *gin.H) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := "https://https://hot-basically-goldfish.ngrok-free.app//backend/notify"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("failed to send request:", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		return err
	}

	log.Println("response:", string(responseBody))
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// Returns new context with specified init data.
func withInitData(ctx context.Context, initData initdata.InitData) context.Context {
	return context.WithValue(ctx, _initDataKey, initData)
}

// Returns the init data from the specified context.
func ctxInitData(ctx context.Context) (initdata.InitData, bool) {
	initData, ok := ctx.Value(_initDataKey).(initdata.InitData)
	return initData, ok
}

// Middleware which authorizes the external client.
func authMiddleware(token string) gin.HandlerFunc {
	return func(context *gin.Context) {
		// We expect passing init data in the Authorization header in the following format:
		// <auth-type> <auth-data>
		// <auth-type> must be "tma", and <auth-data> is Telegram Mini Apps init data.
		authParts := strings.Split(context.GetHeader("authorization"), " ")
		if len(authParts) != 2 {
			context.AbortWithStatusJSON(401, map[string]any{
				"message": "Unauthorized",
			})
			return
		}

		authType := authParts[0]
		authData := authParts[1]

		switch authType {
		case "tma":
			// Validate init data. We consider init data sign valid for 1 hour from their
			// creation moment.
			if err := initdata.Validate(authData, token, time.Hour); err != nil {
				context.AbortWithStatusJSON(401, map[string]any{
					"message": err.Error(),
				})
				return
			}

			// Parse init data. We will surely need it in the future.
			initData, err := initdata.Parse(authData)
			if err != nil {
				context.AbortWithStatusJSON(500, map[string]any{
					"message": err.Error(),
				})
				return
			}

			context.Request = context.Request.WithContext(
				withInitData(context.Request.Context(), initData),
			)
		}
	}
}
