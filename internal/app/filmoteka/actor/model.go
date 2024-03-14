package actor

import "time"

type Actor struct {
	ID        string
	Sex       string
	BirthDate *time.Time
}
