package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/store"
)

type Worker struct {
	ID    int
	Stats *model.Statistics
	Store *store.CounterStore[model.Event]
	WG    *sync.WaitGroup
	Stop  <-chan struct{}
}

func (w *Worker) Start(eventChan <-chan model.Event) {
	defer w.WG.Done()

	for {
		select {
		case <-w.Stop:
			fmt.Printf("Worker %d stopped work\n", w.ID)
			return
		case event, exists := <-eventChan:
			if !exists {
				fmt.Printf("Worker %d stopped work\n", w.ID)
				return
			}

			fmt.Printf("Worker %d in progress on event %s\n", w.ID, event.Type)

			time.Sleep(5 * time.Second)

			w.Store.Mu.Lock()
			w.Stats.Processed++
			w.Store.Inc(event)
			w.Store.Mu.Unlock()

			fmt.Printf("Worker %d finished event %s\n", w.ID, event.Type)
		}
	}
}
