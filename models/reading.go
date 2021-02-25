package models

import "time"

type Reading struct {
	ID        string	`json:"id"`
	Type      string	`json:"type"`
	Value     float64	`json:"value"`
	Alert     bool		`json:"alert"`
	Timestamp time.Time	`json:"timestamp"`
}