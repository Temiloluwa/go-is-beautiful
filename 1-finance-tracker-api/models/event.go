package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
	Tags        []string // Mutable field (slice)
}

var events = []Event{}

// Fix: Use a pointer receiver to modify the actual Event instance
func (e *Event) Save() {
	e.ID = len(events) + 1
	e.UserId = len(events)
	e.AddTag(fmt.Sprintf("event %d", len(events))) // Convert len(events) to a string properly
	events = append(events, *e)                    // Store the dereferenced event
}

func (e *Event) AddTag(tag string) {
	e.Tags = append(e.Tags, tag) // This now modifies the same instance
}

func GetAllEvents() []Event {
	return events
}
