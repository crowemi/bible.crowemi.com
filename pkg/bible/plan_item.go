package bible

import "time"

type PlanItem struct {
	ID        string
	PlanID    string    `firestore:"plan_id,omitempty"`
	BookID    string    `firestore:"book_id,omitempty"`
	ChapterID string    `firestore:"chapter_id,omitempty"`
	ReadDate  string    `firestore:"read_date,omitempty"`
	Summary   string    `firestore:"summary,omitempty"`
	Query     string    `firestore:"query,omitempty"`
	CreatedAt time.Time `firestore:"created_at,omitempty"`
	UpdatedAt time.Time `firestore:"updated_at,omitempty"`
}
