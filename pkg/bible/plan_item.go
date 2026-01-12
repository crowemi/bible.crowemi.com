package bible

type PlanItem struct {
	DocumentID string `firestore:"document_id,omitempty"`
	Passage    string `firestore:"passage,omitempty"`
	Date       string `firestore:"date,omitempty"`
	Plan       string `firestore:"plan,omitempty"`
	Summary    string `firestore:"summary,omitempty"`
	Link       string `firestore:"link,omitempty"`
}
