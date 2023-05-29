package model

type EndpointStats struct {
	Endpoint     string `db:"endpoint" json:"endpoint"`
	RequestCount int64  `db:"request_count" json:"request_count"`
}