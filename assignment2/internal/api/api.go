package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CoffeeSi/golangAITU/assignment2/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/queue"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/store"
)

type Handler struct {
	Store *store.Store
	Queue *queue.Queue[*model.Task]
	Stop  <-chan struct{}
}

func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		h.Store.Mu.Lock()
		var tasksList []model.Task
		for _, task := range h.Store.Tasks {
			tasksList = append(tasksList, *task)
		}
		h.Store.Mu.Unlock()

		json.NewEncoder(w).Encode(tasksList)
	case http.MethodPost:
		var postData model.PostData
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		select {
		case <-h.Stop:
			http.Error(w, "server does not accept new requests", http.StatusServiceUnavailable)
			return
		default:
		}

		h.Store.Mu.Lock()
		h.Store.NextID++
		id := strconv.Itoa(h.Store.NextID)
		task := model.Task{
			ID:     id,
			Status: "PENDING",
		}
		h.Store.Tasks[id] = &task
		h.Store.Stats.Submitted++
		h.Store.Mu.Unlock()

		h.Queue.Enqueue(&task)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(model.IDResponse{ID: id})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskID := r.PathValue("id")
	switch r.Method {
	case http.MethodGet:
		h.Store.Mu.Lock()
		_, exists := h.Store.Tasks[taskID]
		h.Store.Mu.Unlock()

		if !exists {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(h.Store.Tasks[taskID])

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Store.Mu.Lock()
		c := h.Store.Stats
		h.Store.Mu.Unlock()
		json.NewEncoder(w).Encode(c)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
