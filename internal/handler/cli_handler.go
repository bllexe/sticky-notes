package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bllexe/sticky-notes/internal/model"
	"github.com/bllexe/sticky-notes/internal/service"
)

type CLIHandler struct {
	noteService *service.NoteService
	reader      *bufio.Reader
}

func NewCLIHandler(noteService *service.NoteService) *CLIHandler {
	return &CLIHandler{
		noteService: noteService,
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (h *CLIHandler) Start() {
	fmt.Println("Welcome to Sticky Notes Application!")
	fmt.Println("===================================")

	for {
		h.printMenu()
		choice := h.readInput("Enter your choice: ")

		switch choice {
		case "1":
			h.createNote()
		case "2":
			h.listNotes()
		case "3":
			h.updateNote()
		case "4":
			h.deleteNote()
		case "5":
			h.searchNotes()
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (h *CLIHandler) printMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Create new note")
	fmt.Println("2. List all notes")
	fmt.Println("3. Update note")
	fmt.Println("4. Delete note")
	fmt.Println("5. Search notes")
	fmt.Println("6. Exit")
}

func (h *CLIHandler) readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := h.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (h *CLIHandler) createNote() {
	content := h.readInput("Enter note content: ")
	color := h.selectColor()

	note, err := h.noteService.CreateNote(content, color)
	if err != nil {
		fmt.Printf("Error creating note: %v\n", err)
		return
	}

	fmt.Printf("Note created successfully with ID: %s\n", note.ID)
}

func (h *CLIHandler) listNotes() {
	notes, err := h.noteService.GetAllNotes()
	if err != nil {
		fmt.Printf("Error getting notes: %v\n", err)
		return
	}

	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	fmt.Println("\nYour Notes:")
	for _, note := range notes {
		h.printNote(note)
	}
}

func (h *CLIHandler) updateNote() {
	id := h.readInput("Enter note ID to update: ")

	note, err := h.noteService.GetNote(id)
	if err != nil {
		fmt.Printf("Error finding note: %v\n", err)
		return
	}

	fmt.Printf("Current content: %s\n", note.Content)
	content := h.readInput("Enter new content (press Enter to keep current): ")
	if content == "" {
		content = note.Content
	}

	color := h.selectColor()

	updatedNote, err := h.noteService.UpdateNote(id, content, color)
	if err != nil {
		fmt.Printf("Error updating note: %v\n", err)
		return
	}

	fmt.Println("Note updated successfully!")
	h.printNote(updatedNote)
}

func (h *CLIHandler) deleteNote() {
	id := h.readInput("Enter note ID to delete: ")

	err := h.noteService.DeleteNote(id)
	if err != nil {
		fmt.Printf("Error deleting note: %v\n", err)
		return
	}

	fmt.Println("Note deleted successfully!")
}

func (h *CLIHandler) searchNotes() {
	query := h.readInput("Enter search query: ")

	notes, err := h.noteService.SearchNotes(query)
	if err != nil {
		fmt.Printf("Error searching notes: %v\n", err)
		return
	}

	if len(notes) == 0 {
		fmt.Println("No matching notes found.")
		return
	}

	fmt.Printf("\nFound %d matching notes:\n", len(notes))
	for _, note := range notes {
		h.printNote(note)
	}
}

func (h *CLIHandler) selectColor() model.Color {
	fmt.Println("\nAvailable colors:")
	fmt.Println("1. Yellow")
	fmt.Println("2. Blue")
	fmt.Println("3. Green")
	fmt.Println("4. Pink")
	fmt.Println("5. Orange")

	choice := h.readInput("Select color (1-5) [default: Yellow]: ")

	switch choice {
	case "2":
		return model.Blue
	case "3":
		return model.Green
	case "4":
		return model.Pink
	case "5":
		return model.Orange
	default:
		return model.Yellow
	}
}

func (h *CLIHandler) printNote(note *model.Note) {
	fmt.Printf("\nID: %s\n", note.ID)
	fmt.Printf("Content: %s\n", note.Content)
	fmt.Printf("Color: %s\n", note.Color)
	fmt.Printf("Created: %s\n", note.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", note.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("------------------------")
}
