package model

import (
	"testing"
	"time"
)

func TestColorValidation(t *testing.T) {

	tests := []struct {
		name  string
		color Color
		want  bool
	}{
		{
			name:  "Valid Yellow Color",
			color: Yellow,
			want:  true,
		},
		{
			name:  "Valid Blue Color",
			color: Blue,
			want:  true,
		},
		{
			name:  "Valid Green Color",
			color: Green,
			want:  true,
		},
		{
			name:  "Valid Pink Color",
			color: Pink,
			want:  true,
		},
		{
			name:  "Valid Orange Color",
			color: Orange,
			want:  true,
		},
		{
			name:  "Invalid Color",
			color: "purple",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := false
			switch tt.color {
			case Yellow, Blue, Green, Pink, Orange:
				isValid = true
			}

			if isValid != tt.want {
				t.Errorf("Color validation failed for %s, got: %v, want: %v", tt.color, isValid, tt.want)
			}
		})
	}
}

func TestNoteCreation(t *testing.T) {
	note := &Note{
		ID:        "test-id",
		Content:   "test-content",
		Color:     "test-color",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if note.ID != "test-id" {
		t.Errorf("Expected note ID to be 'test-id', got: %s", note.ID)
	}
	if note.Content != "test-content" {
		t.Errorf("Expected note content to be 'test-content', got: %s", note.Content)
	}
	if note.Color != "test-color" {
		t.Errorf("Expected note color to be 'test-color', got: %s", note.Color)
	}
	if note.CreatedAt.IsZero() {
		t.Errorf("Expected note created at to be non-zero, got: %s", note.CreatedAt)
	}
	if note.UpdatedAt.IsZero() {
		t.Errorf("Expected note updated at to be non-zero, got: %s", note.UpdatedAt)
	}
}
