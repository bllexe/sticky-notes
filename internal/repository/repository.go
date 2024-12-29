package repository

import (
	"github.com/bllexe/sticky-notes/internal/model"
)

type NoteRepository interface {
	Save(note *model.Note) error
	Update(note *model.Note) error
	Delete(id string) error
	GetById(id string) (*model.Note, error)
	GetAll() ([]*model.Note, error)
	Search(query string) ([]*model.Note, error)
}
