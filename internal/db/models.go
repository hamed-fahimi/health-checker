package db

import "time"

type MonitoredService struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	URL     string
	Enabled bool
}

type CheckResult struct {
	ID           uint `gorm:"primaryKey"`
	URL          string
	ServiceName  string
	Status       string
	StatusCode   int
	ResponseTime string
	CheckedAt    time.Time
}
