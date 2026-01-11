package bible

type Book struct {
	BookID   string        `firestore:"book_id,omitempty"`
	Name     string        `firestore:"name,omitempty"`
	Number   int           `firestore:"number,omitempty"`
	Chapters []BookChapter `firestore:"chapters,omitempty"`
	Verses   int           `firestore:"verses,omitempty"`
}

type BookChapter struct {
	ChapterID string `firestore:"chapter_id,omitempty"`
	Number    int    `firestore:"number,omitempty"`
	Verses    int    `firestore:"verses,omitempty"`
}
