package repository

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/bllexe/sticky-notes/internal/model"
)

func setupTestRepo(t *testing.T) (*FileRepository, string) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "sticky-notes-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	repo, err := NewFileRepository(tempDir)
	if err != nil {
		t.Fatalf("Failed to create repository: %v", err)
	}

	return repo, tempDir
}

func cleanupTestRepo(tempDir string) {
	os.RemoveAll(tempDir)
}

func createTestNote() *model.Note {
	return &model.Note{
		ID:        "test-note-id",
		Content:   "test content",
		Color:     model.Yellow,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestSave(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	note := createTestNote()

	// Test saving a note
	err := repo.Save(note)
	if err != nil {
		t.Errorf("Failed to save note: %v", err)
	}

	// Verify file exists
	expectedPath := filepath.Join(tempDir, note.ID+".json")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Note file was not created at %s", expectedPath)
	}
}

func TestGetByID(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	note := createTestNote()

	// Save a note first
	err := repo.Save(note)
	if err != nil {
		t.Fatalf("Failed to save note: %v", err)
	}

	// Test getting the note
	retrieved, err := repo.GetByID(note.ID)
	if err != nil {
		t.Errorf("Failed to get note by ID: %v", err)
	}

	if retrieved.ID != note.ID {
		t.Errorf("Retrieved note ID mismatch, got: %s, want: %s", retrieved.ID, note.ID)
	}

	if retrieved.Content != note.Content {
		t.Errorf("Retrieved note content mismatch, got: %s, want: %s", retrieved.Content, note.Content)
	}

	// Test getting non-existent note
	_, err = repo.GetByID("non-existent-id")
	if err == nil {
		t.Error("Expected error when getting non-existent note, got nil")
	}
}

func TestUpdate(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	note := createTestNote()

	// Save initial note
	err := repo.Save(note)
	if err != nil {
		t.Fatalf("Failed to save note: %v", err)
	}

	// Update note
	note.Content = "updated content"
	note.Color = model.Blue
	err = repo.Update(note)
	if err != nil {
		t.Errorf("Failed to update note: %v", err)
	}

	// Verify update
	updated, err := repo.GetByID(note.ID)
	if err != nil {
		t.Fatalf("Failed to get updated note: %v", err)
	}

	if updated.Content != "updated content" {
		t.Errorf("Note content was not updated, got: %s, want: updated content", updated.Content)
	}

	if updated.Color != model.Blue {
		t.Errorf("Note color was not updated, got: %s, want: %s", updated.Color, model.Blue)
	}
}

func TestDelete(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	note := createTestNote()

	// Save a note first
	err := repo.Save(note)
	if err != nil {
		t.Fatalf("Failed to save note: %v", err)
	}

	// Test deleting the note
	err = repo.Delete(note.ID)
	if err != nil {
		t.Errorf("Failed to delete note: %v", err)
	}

	// Verify note is deleted
	_, err = repo.GetByID(note.ID)
	if err == nil {
		t.Error("Expected error when getting deleted note, got nil")
	}
}

func TestGetAll(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	// Create multiple notes
	notes := []*model.Note{
		{
			ID:      "note1",
			Content: "content1",
			Color:   model.Yellow,
		},
		{
			ID:      "note2",
			Content: "content2",
			Color:   model.Blue,
		},
	}

	// Save all notes
	for _, note := range notes {
		err := repo.Save(note)
		if err != nil {
			t.Fatalf("Failed to save note: %v", err)
		}
	}

	// Test getting all notes
	retrieved, err := repo.GetAll()
	if err != nil {
		t.Errorf("Failed to get all notes: %v", err)
	}

	if len(retrieved) != len(notes) {
		t.Errorf("Retrieved notes count mismatch, got: %d, want: %d", len(retrieved), len(notes))
	}
}

func TestSearch(t *testing.T) {
	repo, tempDir := setupTestRepo(t)
	defer cleanupTestRepo(tempDir)

	// Create notes with different content
	notes := []*model.Note{
		{
			ID:      "note1",
			Content: "apple pie recipe",
			Color:   model.Yellow,
		},
		{
			ID:      "note2",
			Content: "shopping list",
			Color:   model.Blue,
		},
		{
			ID:      "note3",
			Content: "apple shopping",
			Color:   model.Green,
		},
	}

	// Save all notes
	for _, note := range notes {
		err := repo.Save(note)
		if err != nil {
			t.Fatalf("Failed to save note: %v", err)
		}
	}

	// Test searching notes
	results, err := repo.Search("apple")
	if err != nil {
		t.Errorf("Failed to search notes: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("Search results count mismatch, got: %d, want: 2", len(results))
	}

	// Test search with no results
	results, err = repo.Search("banana")
	if err != nil {
		t.Errorf("Failed to search notes: %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected no results for 'banana' search, got: %d results", len(results))
	}
}
