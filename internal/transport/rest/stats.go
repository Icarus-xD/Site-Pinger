package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getEndpointsStats(ctx *gin.Context) {
	stats, err := h.statsService.GetEndpointsStats()
	if err != nil {
		log.Println("getEndpointsStats() error:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, stats)
}