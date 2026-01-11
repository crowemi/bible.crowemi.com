package bible

type Plan struct {
	PlanID string     `firestore:"plan_id,omitempty"`
	Name   string     `firestore:"name,omitempty"`
	Items  []PlanItem `firestore:"items,omitempty"`
}
