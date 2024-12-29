package service

import (
	"fmt"
	"time"

	"github.com/bllexe/sticky-notes/internal/model"
	"github.com/bllexe/sticky-notes/internal/repository"
	"github.com/google/uuid"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) CreateNote(content string, color model.Color) (*model.Note, error) {
	note := &model.Note{
		ID:        uuid.New().String(),
		Content:   content,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.validateNote(note); err != nil {
		return nil, err
	}

	if err := s.repo.Save(note); err != nil {
		return nil, fmt.Errorf("failed to save note: %w", err)
	}

	return note, nil
}

func (s *NoteService) UpdateNote(id string, content string, color model.Color) (*model.Note, error) {
	note, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get note: %w", err)
	}

	note.Content = content
	note.Color = color
	note.UpdatedAt = time.Now()

	if err := s.validateNote(note); err != nil {
		return nil, err
	}

	if err := s.repo.Update(note); err != nil {
		return nil, fmt.Errorf("failed to update note: %w", err)
	}

	return note, nil
}

func (s *NoteService) DeleteNote(id string) error {
	return s.repo.Delete(id)
}

func (s *NoteService) GetNote(id string) (*model.Note, error) {
	return s.repo.GetById(id)
}

func (s *NoteService) GetAllNotes() ([]*model.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) SearchNotes(query string) ([]*model.Note, error) {
	return s.repo.Search(query)
}

func (s *NoteService) validateNote(note *model.Note) error {
	if note.Content == "" {
		return fmt.Errorf("note content cannot be empty")
	}

	switch note.Color {
	case model.Yellow, model.Blue, model.Green, model.Pink, model.Orange:
		return nil
	default:
		return fmt.Errorf("invalid note color: %s", note.Color)
	}
}
