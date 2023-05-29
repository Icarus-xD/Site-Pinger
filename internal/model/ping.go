package model

import "github.com/google/uuid"

type SitePing struct {
	ID   uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Site string      `gorm:"not null" json:"site"`
	Ping int64       `gorm:"not null" json:"ping"`
}

func (SitePing) TableName() string {
	return "sites_ping"
}