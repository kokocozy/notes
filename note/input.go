package note

import (
	"notes-api/user"
	"time"
)

type NewNoteInput struct {
	Title  string `json:"title" binding:"required"`
	Detail string `json:"detail" binding:"required"`
	Status bool   `json:"status" binding:"required"`
}

type UpdateNoteInput struct {
	Title     string `json:"title" binding:"required"`
	Detail    string `json:"detail" binding:"required"`
	Status    bool
	UpdatedAt time.Time
	User      user.User
}

type GetNoteIdInput struct {
	Id int `uri:"id" binding:"required"`
}
