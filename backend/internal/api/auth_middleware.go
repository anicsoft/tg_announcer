package api

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

const (
	InitDataKey = "init-data"
	GuestKey    = "is-guest"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			authorizeGuest(ctx)
			return
		}

		authType, authData := authParts[0], authParts[1]

		if authType != "tma" {
			authorizeGuest(ctx)
			return
		}

		if err := initdata.Validate(authData, os.Getenv("TELEGRAM_BOT_TOKEN"), time.Hour); err != nil {
			authorizeGuest(ctx)
			return
		}

		initData, err := initdata.Parse(authData)
		if err != nil {
			authorizeGuest(ctx)
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), InitDataKey, initData)
		newCtx = context.WithValue(newCtx, GuestKey, false)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}

func authorizeGuest(ctx *gin.Context) {
	newCtx := context.WithValue(ctx.Request.Context(), InitDataKey, initdata.InitData{}) // Setting empty InitData
	newCtx = context.WithValue(newCtx, GuestKey, true)
	ctx.Request = ctx.Request.WithContext(newCtx)
	ctx.Next()
}

func GetInitData(ctx context.Context) (initdata.InitData, bool, bool) {
	initData, initDataExists := ctx.Value(InitDataKey).(initdata.InitData)
	isGuest, isGuestExists := ctx.Value(GuestKey).(bool)

	if !initDataExists {
		return initdata.InitData{}, false, isGuestExists && isGuest
	}

	return initData, true, isGuestExists && isGuest
}
