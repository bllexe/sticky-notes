package repository

import (
	"fmt"
	"os"
	"sync"
)

type FileRepository struct {
	dataDir string
	mutex   sync.Mutex
}

func NewFileRepository(dataDir string) (*FileRepository, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {

		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}
	return &FileRepository{dataDir: dataDir}, nil
}
