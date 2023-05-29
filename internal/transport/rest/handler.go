package rest

import (
	"github.com/Icarus-xD/SitePinger/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Ping interface {
	GetById(id uuid.UUID) (model.SitePing, error)
	GetMin() ([]model.SitePing, error)
	GetMax() ([]model.SitePing, error)
}

type Stats interface {
	SaveStat(endpoint string) error
	GetEndpointsStats() ([]model.EndpointStats, error)
}

type Handler struct {
	pingService Ping
	statsService Stats
}

func NewHandler(p Ping, s Stats) *Handler {
	return &Handler{
		pingService: p,
		statsService: s,
	}
}

func (h *Handler) InitRouter(r *gin.Engine) {
	r.Use(h.statsMiddleware())

	pingRoutes := r.Group("ping")
	{
		pingRoutes.GET("/:id", h.getPingById)
		pingRoutes.GET("/min", h.getMinPing)
		pingRoutes.GET("/max", h.getMaxPing)
	}

	statsRoutes := r.Group("stats")
	statsRoutes.Use(h.authMiddleware())
	{
		statsRoutes.GET("/endpoints", h.getEndpointsStats)
	}
}