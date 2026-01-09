package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/store"
)

type Handler struct {
	Store     *store.CounterStore[model.Event]
	Stats     *model.Statistics
	Stop      <-chan struct{}
	EventChan chan model.Event
}

func (h *Handler) EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	select {
	case <-h.Stop:
		http.Error(w, "server does not accept new requests", http.StatusServiceUnavailable)
		return
	default:
	}

	switch r.Method {
	case http.MethodPost:
		var postData model.Event
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		if postData.Type == "" {
			http.Error(w, "Missing or empty type", http.StatusBadRequest)
			return
		}

		h.Store.Mu.Lock()
		h.Stats.Received++
		h.Store.Mu.Unlock()
		h.EventChan <- postData

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"message": "Event accepted"})
		fmt.Printf("Event accepted: %s\n", postData.Type)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) CountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	select {
	case <-h.Stop:
		http.Error(w, "server does not accept new requests", http.StatusServiceUnavailable)
		return
	default:
	}

	switch r.Method {
	case http.MethodGet:
		h.Store.Mu.Lock()
		counts := h.Store.GetAll()
		h.Store.Mu.Unlock()

		result := make(map[string]int)
		for event, count := range counts {
			result[event.Type] += count
		}

		json.NewEncoder(w).Encode(result)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		select {
		case <-h.Stop:
			http.Error(w, "server does not accept new requests", http.StatusServiceUnavailable)
			return
		default:
		}

		h.Store.Mu.Lock()
		stats := *h.Stats
		stats.Queued = uint64(len(h.EventChan))
		stats.UniqueTypes = uint64(len(h.Store.GetAll()))
		h.Store.Mu.Unlock()

		json.NewEncoder(w).Encode(map[string]uint64{
			"received":     stats.Received,
			"processed":    stats.Processed,
			"queued":       stats.Queued,
			"unique_types": stats.UniqueTypes})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
