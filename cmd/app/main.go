package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bllexe/sticky-notes/internal/handler"
	"github.com/bllexe/sticky-notes/internal/repository"
	"github.com/bllexe/sticky-notes/internal/service"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Setup data directory
	dataDir := filepath.Join(currentDir, "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Initialize repository
	repo, err := repository.NewFileRepository(dataDir)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	// Initialize service
	noteService := service.NewNoteService(repo)

	// Initialize and start CLI handler
	cli := handler.NewCLIHandler(noteService)
	cli.Start()
}
