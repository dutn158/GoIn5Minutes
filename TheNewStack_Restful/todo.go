package main

import (
	"time"
)

// Todo struct
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos slices of Todo
type Todos []Todo
