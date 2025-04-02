package domain

import "time"

type File struct {
	Name      string
	Content   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
