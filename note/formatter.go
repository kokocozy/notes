package note

import "time"

type NoteFormatter struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatNote(note Note) *NoteFormatter {
	formatter := new(NoteFormatter)
	formatter.Id = note.Id
	formatter.Title = note.Title
	formatter.Detail = note.Detail
	formatter.Status = note.Status
	formatter.CreatedAt = note.CreatedAt
	formatter.UpdatedAt = note.UpdatedAt

	return formatter
}

func FormatNotes(notes []Note) []NoteFormatter {
	notesFormatter := []NoteFormatter{}

	for _, note := range notes {
		noteFormatter := FormatNote(note)
		notesFormatter = append(notesFormatter, *noteFormatter)
	}

	return notesFormatter
}

// contoh pointer bool
type FtterNote struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Status    *bool     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NoteFtter(note Note) FtterNote {
	formatter := FtterNote{
		Id:        note.Id,
		Title:     note.Title,
		Detail:    note.Detail,
		Status:    &note.Status,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}

	return formatter
}
