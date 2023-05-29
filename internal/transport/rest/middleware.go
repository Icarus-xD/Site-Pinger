package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/Icarus-xD/SitePinger/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type CtxValue int

const (
	ctxAuth CtxValue = iota
)

func (h *Handler) statsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := h.statsService.SaveStat(ctx.FullPath())
		if err != nil {
			log.Println("inserting into clickhouse endpoints_stats:", err)
		}
	}
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helpers.GetTokenFromRequest(c)
		if err != nil {
			log.Println("AuthMiddleware", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c.Request.Context(), ctxAuth, token)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}