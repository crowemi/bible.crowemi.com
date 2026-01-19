package bible

import "time"

type Plan struct {
	ID          string
	Name        string    `firestore:"name,omitempty"`
	Description string    `firestore:"description,omitempty"`
	CreatedAt   time.Time `firestore:"created_at,omitempty"`
	UpdatedAt   time.Time `firestore:"updated_at,omitempty"`
}
