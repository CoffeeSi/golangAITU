package store

import (
	"sync"

	"github.com/CoffeeSi/golangAITU/assignment2/internal/model"
)

type Store struct {
	Mu     sync.Mutex
	Tasks  map[string]*model.Task
	Stats  model.Stats
	NextID int
	WG     sync.WaitGroup
}

func New() *Store {
	return &Store{
		Tasks: make(map[string]*model.Task),
	}
}
