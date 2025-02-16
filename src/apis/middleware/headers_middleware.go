package middleware

import (
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/configs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddHeadersSecurity(cfg *configs.Configs) gin.HandlerFunc {
	origin := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	return func(ctx *gin.Context) {
		if ctx.Request.Host != origin {
			helpers.SendResult(&helpers.Result{Ok: false, Status: 400, Message: "هاست نامعتبر", Data: nil}, ctx)
			ctx.Abort()
			return
		}
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		ctx.Header("Referrer-Policy", "strict-origin")
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		ctx.Next()
	}

}
