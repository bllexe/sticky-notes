package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/bllexe/sticky-notes/internal/model"
)

type FileRepository struct {
	dataDir string
	mutex   sync.RWMutex
}

func NewFileRepository(dataDir string) (*FileRepository, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}
	return &FileRepository{
		dataDir: dataDir,
	}, nil
}

func (r *FileRepository) Save(note *model.Note) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	filename := filepath.Join(r.dataDir, note.ID+".json")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create note file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(note); err != nil {
		return fmt.Errorf("failed to encode note: %w", err)
	}
	return nil
}

func (r *FileRepository) Update(note *model.Note) error {
	return r.Save(note)
}

func (r *FileRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	filename := filepath.Join(r.dataDir, id+".json")
	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}
	return nil
}

func (r *FileRepository) GetByID(id string) (*model.Note, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	filename := filepath.Join(r.dataDir, id+".json")
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open note file: %w", err)
	}
	defer file.Close()

	var note model.Note
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&note); err != nil {
		return nil, fmt.Errorf("failed to decode note: %w", err)
	}
	return &note, nil
}

func (r *FileRepository) GetAll() ([]*model.Note, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	files, err := os.ReadDir(r.dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read data directory: %w", err)
	}

	var notes []*model.Note
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		note, err := r.GetByID(strings.TrimSuffix(file.Name(), ".json"))
		if err != nil {
			continue
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *FileRepository) Search(query string) ([]*model.Note, error) {
	notes, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	query = strings.ToLower(query)
	var results []*model.Note
	for _, note := range notes {
		if strings.Contains(strings.ToLower(note.Content), query) {
			results = append(results, note)
		}
	}
	return results, nil
}
