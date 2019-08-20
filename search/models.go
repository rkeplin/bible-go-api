package search

type Book struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Testament string `json:"testament"`
}

type Text struct {
	ID        int64  `json:"id"`
	Book      Book   `json:"book"`
	ChapterID int64  `json:"chapterId"`
	VerseID   int64  `json:"verseId"`
	Verse     string `json:"verse"`
}

type TextCollection struct {
	Total int64  `json:"total"`
	Items []Text `json:"items"`
}

type ESHighlight struct {
	Verse []string `json:"verse"`
}

type ESSource struct {
	ID        int64  `json:"id"`
	Testament string `json:"testament"`
	BookID    int64  `json:"bookId"`
	BookName  string `json:"bookName"`
	ChapterID int64  `json:"chapterId"`
	VerseID   int64  `json:"verseId"`
	Verse     string `json:"verse"`
}

type ESInnerHit struct {
	Source    ESSource    `json:"_source"`
	Highlight ESHighlight `json:"highlight"`
}

type ESTotal struct {
	Value int64 `json:"value"`
}

type ESHits struct {
	Total ESTotal      `json:"total"`
	Hits  []ESInnerHit `json:"hits"`
}

type ESResult struct {
	Hits ESHits `json:"hits"`
}
