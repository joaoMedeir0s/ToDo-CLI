package main

import "time"

type Task struct {
	// The name of the variables have to be capitalized in order to be exported
	//  and accessible from other packages.
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}
