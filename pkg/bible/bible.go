package bible

import "time"

type Book struct {
	ID           string    `firestore:"id"`
	Name         string    `firestore:"name"`
	Osis         string    `firestore:"osis,omitempty"`
	ShortName    string    `firestore:"short_name,omitempty"`
	Number       string    `firestore:"number,omitempty"`
	BookDivision string    `firestore:"book_div,omitempty"`
	Testament    string    `firestore:"testament,omitempty"`
	Chapters     []Chapter `firestore:"chapters,omitempty"`
	ChapterCount int       `firestore:"chapter_count,omitempty"`
	Summary      string    `firestore:"summary,omitempty"`
	CreatedAt    time.Time `firestore:"created_at,omitempty"`
	UpdatedAt    time.Time `firestore:"updated_at,omitempty"`
	Slug         string    `firestore:"slug,omitempty"`
	YearWritten  string    `firestore:"year_written,omitempty"`
	PlaceWritten string    `firestore:"place_written,omitempty"`
	VerseCount   int       `firestore:"verse_count,omitempty"`
	Writers      string    `firestore:"writers,omitempty"`
	PeopleCount  int       `firestore:"people_count,omitempty"`
	PlaceCount   int       `firestore:"place_count,omitempty"`
}

type Chapter struct {
	ID          string    `firestore:"id"`
	Osis        string    `firestore:"osis,omitempty"`
	Number      string    `firestore:"number,omitempty"`
	Verses      []Verse   `firestore:"verses,omitempty"`
	Summary     string    `firestore:"summary,omitempty"`
	Writer      string    `firestore:"writer,omitempty"`
	Slug        string    `firestore:"slug,omitempty"`
	PeopleCount int       `firestore:"people_count,omitempty"`
	PlaceCount  int       `firestore:"places_count,omitempty"`
	WriterCount int       `firestore:"writer_count,omitempty"`
	CreatedAt   time.Time `firestore:"created_at,omitempty"`
	UpdatedAt   time.Time `firestore:"updated_at,omitempty"`
}

type Verse struct {
	Number       int       `firestore:"number,omitempty"`
	Summary      string    `firestore:"summary,omitempty"`
	CreatedAt    time.Time `firestore:"created_at,omitempty"`
	UpdatedAt    time.Time `firestore:"updated_at,omitempty"`
	Osis         string    `firestore:"osis,omitempty"`
	VerseID      string    `firestore:"verse_id,omitempty"`
	Book         string    `firestore:"book,omitempty"`
	Chapter      string    `firestore:"chapter,omitempty"`
	VerseNum     int       `firestore:"verse_num,omitempty"`
	VerseText    string    `firestore:"verse_text,omitempty"`
	RichText     string    `firestore:"rich_text,omitempty"`
	MdText       string    `firestore:"md_text,omitempty"`
	People       []string  `firestore:"people,omitempty"`
	PeopleCount  int       `firestore:"people_count,omitempty"`
	Places       []string  `firestore:"places,omitempty"`
	PlacesCount  int       `firestore:"places_count,omitempty"`
	YearNum      int       `firestore:"year_num,omitempty"`
	PeopleGroups []string  `firestore:"people_groups,omitempty"`
	Event        string    `firestore:"event,omitempty"`
}
