// Package models ..
package models

import "time"

type EventDefinitions struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Event       string `json:"event" db:"event"`
	Description string `json:"description,omitempty" db:"description"`
}

type Events struct {
	ID         int64          `json:"id,omitempty" db:"id"`
	Event      string         `json:"event" db:"event"`
	DistinctID string         `json:"distinct_id,omitempty" db:"distinct_id"`
	Timestamp  time.Time      `json:"timestamp" db:"timestamp"`
	Properties map[string]any `json:"properties,omitempty" db:"properties"`
}
