package simple_action

import "time"

type Task struct {
	Name        string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	DoneAt      *time.Time
}
