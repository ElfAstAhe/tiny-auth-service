package dto

import (
	"time"
)

type LoginAttemptEventDTO struct {
	NodeName  string    `json:"node_name"`
	EventDate time.Time `json:"event_date"`
	RequestID string    `json:"request_id"`
	TraceID   string    `json:"trace_id"`
	IP        string    `json:"ip"`
	Username  string    `json:"username"`
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
}
