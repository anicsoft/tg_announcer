package api

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

const (
	InitDataKey = "init-data"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			StatusUnauthorizedWithAbort(ctx, ErrUnauthorized)
			return
		}

		authType, authData := authParts[0], authParts[1]

		if authType != "tma" {
			StatusUnauthorizedWithAbort(ctx, ErrUnauthorized)
			return
		}

		if err := initdata.Validate(authData, os.Getenv("TELEGRAM_BOT_TOKEN"), time.Hour); err != nil {
			StatusUnauthorizedWithAbort(ctx, errors.Join(err, ErrUnauthorized))
			return
		}

		initData, err := initdata.Parse(authData)
		if err != nil {
			StatusUnauthorizedWithAbort(ctx, errors.Join(err, ErrUnauthorized))
			return
		}

		ctx.Set(InitDataKey, initData)
		ctx.Next()
	}
}

func GetInitData(ctx *gin.Context) (initdata.InitData, bool) {
	initData, exists := ctx.Get(InitDataKey)
	if !exists {
		return initdata.InitData{}, false
	}

	parsedData, ok := initData.(initdata.InitData)
	return parsedData, ok
}
