package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) getPingById(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		log.Println("getSitePing() error:", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ping, err := h.pingService.GetById(parsedID)
	if err != nil {
		log.Println("getSitePing() error:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, ping)
}

func (h *Handler) getMinPing(ctx *gin.Context) {
	pings, err := h.pingService.GetMin()
	if err != nil {
		log.Println("getMinPing() error:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, pings)
}

func (h *Handler) getMaxPing(ctx *gin.Context) {
	pings, err := h.pingService.GetMax()
	if err != nil {
		log.Println("getMinPing() error:", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, pings)
}