package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/bllexe/sticky-notes/internal/model"
)

// MockRepository is a mock implementation of repository.NoteRepository
type MockRepository struct {
	notes map[string]*model.Note
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		notes: make(map[string]*model.Note),
	}
}

func (r *MockRepository) Save(note *model.Note) error {
	r.notes[note.ID] = note
	return nil
}

func (r *MockRepository) Update(note *model.Note) error {
	if _, exists := r.notes[note.ID]; !exists {
		return fmt.Errorf("note not found")
	}
	r.notes[note.ID] = note
	return nil
}

func (r *MockRepository) Delete(id string) error {
	delete(r.notes, id)
	return nil
}

func (r *MockRepository) GetById(id string) (*model.Note, error) {
	note, exists := r.notes[id]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}
	return note, nil
}

func (r *MockRepository) GetAll() ([]*model.Note, error) {
	var notes []*model.Note
	for _, note := range r.notes {
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *MockRepository) Search(query string) ([]*model.Note, error) {
	return []*model.Note{}, nil // Simplified for testing
}

func TestCreateNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewNoteService(repo)

	tests := []struct {
		name      string
		content   string
		color     model.Color
		wantError bool
	}{
		{
			name:      "Valid Note",
			content:   "Test content",
			color:     model.Yellow,
			wantError: false,
		},
		{
			name:      "Empty Content",
			content:   "",
			color:     model.Yellow,
			wantError: true,
		},
		{
			name:      "Invalid Color",
			content:   "Test content",
			color:     "invalid-color",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := service.CreateNote(tt.content, tt.color)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if note.Content != tt.content {
				t.Errorf("Content mismatch, got: %s, want: %s", note.Content, tt.content)
			}

			if note.Color != tt.color {
				t.Errorf("Color mismatch, got: %s, want: %s", note.Color, tt.color)
			}

			if note.ID == "" {
				t.Error("Expected ID to be set")
			}

			if note.CreatedAt.IsZero() {
				t.Error("Expected CreatedAt to be set")
			}

			if note.UpdatedAt.IsZero() {
				t.Error("Expected UpdatedAt to be set")
			}
		})
	}
}

func TestUpdateNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewNoteService(repo)

	// Create initial note
	initialNote := &model.Note{
		ID:        "test-id",
		Content:   "initial content",
		Color:     model.Yellow,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo.Save(initialNote)

	tests := []struct {
		name      string
		id        string
		content   string
		color     model.Color
		wantError bool
	}{
		{
			name:      "Valid Update",
			id:        "test-id",
			content:   "updated content",
			color:     model.Blue,
			wantError: false,
		},
		{
			name:      "Empty Content",
			id:        "test-id",
			content:   "",
			color:     model.Blue,
			wantError: true,
		},
		{
			name:      "Invalid Color",
			id:        "test-id",
			content:   "content",
			color:     "invalid-color",
			wantError: true,
		},
		{
			name:      "Non-existent Note",
			id:        "non-existent",
			content:   "content",
			color:     model.Yellow,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedNote, err := service.UpdateNote(tt.id, tt.content, tt.color)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if updatedNote.Content != tt.content {
				t.Errorf("Content mismatch, got: %s, want: %s", updatedNote.Content, tt.content)
			}

			if updatedNote.Color != tt.color {
				t.Errorf("Color mismatch, got: %s, want: %s", updatedNote.Color, tt.color)
			}
		})
	}
}

func TestDeleteNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewNoteService(repo)

	// Create a note to delete
	note := &model.Note{
		ID:      "test-id",
		Content: "content",
		Color:   model.Yellow,
	}
	repo.Save(note)

	err := service.DeleteNote(note.ID)
	if err != nil {
		t.Errorf("Failed to delete note: %v", err)
	}

	// Verify note is deleted
	retrieved, _ := repo.GetById(note.ID)
	if retrieved != nil {
		t.Error("Note should have been deleted")
	}
}

func TestValidateNote(t *testing.T) {
	service := NewNoteService(NewMockRepository())

	tests := []struct {
		name      string
		note      *model.Note
		wantError bool
	}{
		{
			name: "Valid Note",
			note: &model.Note{
				Content: "valid content",
				Color:   model.Yellow,
			},
			wantError: false,
		},
		{
			name: "Empty Content",
			note: &model.Note{
				Content: "",
				Color:   model.Yellow,
			},
			wantError: true,
		},
		{
			name: "Invalid Color",
			note: &model.Note{
				Content: "content",
				Color:   "invalid-color",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.validateNote(tt.note)

			if tt.wantError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
