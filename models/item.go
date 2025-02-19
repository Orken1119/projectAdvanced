package models

import "time"

type ItemDetail struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Location      string    `json:"location"`
	Date          time.Time `json:"date" gorm:"default:CURRENT_TIMESTAMP"`
	ClaimedStatus string    `json:"claimed_status"`
}