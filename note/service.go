package note

import "errors"

type Service interface {
	CreateNote(userId int, input NewNoteInput) (Note, error)
	AllNote() ([]Note, error)
	AllNoteByUserId(id int) ([]Note, error)
	FindNote(id int) (Note, error)
	UpdateNote(id int, input UpdateNoteInput) (Note, error)
	DeleteNote(id int, userId int) (Note, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateNote(userId int, input NewNoteInput) (Note, error) {
	note := Note{}
	note.UserId = userId
	note.Title = input.Title
	note.Detail = input.Detail
	note.Status = input.Status

	newNote, err := s.repository.Save(note)
	if err != nil {
		return newNote, err
	}

	return newNote, nil
}

func (s *service) AllNote() ([]Note, error) {
	notes, err := s.repository.All(0)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (s *service) AllNoteByUserId(id int) ([]Note, error) {
	notes, err := s.repository.All(id)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (s *service) FindNote(id int) (Note, error) {
	note, err := s.repository.First(id)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (s *service) UpdateNote(id int, input UpdateNoteInput) (Note, error) {
	note, err := s.repository.First(id)
	if err != nil {
		return note, err
	}

	if note.UserId != input.User.ID {
		return note, errors.New("Note is not yours")
	}

	note.Title = input.Title
	note.Detail = input.Detail
	note.Status = input.Status
	note.UpdatedAt = input.UpdatedAt
	updatedNote, err := s.repository.Update(note)
	if err != nil {
		return updatedNote, err
	}

	return updatedNote, nil
}

func (s *service) DeleteNote(id int, userId int) (Note, error) {
	note, err := s.repository.First(id)
	if err != nil {
		return note, err
	}

	if note.UserId != userId {
		return note, errors.New("Note is not yours")
	}

	deletedNote, err := s.repository.Delete(note)
	if err != nil {
		return deletedNote, err
	}

	return deletedNote, nil
}
