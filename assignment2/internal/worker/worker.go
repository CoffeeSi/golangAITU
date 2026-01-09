package worker

import (
	"fmt"
	"time"

	"github.com/CoffeeSi/golangAITU/assignment2/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/queue"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/store"
)

func Worker(workerID int, taskQueue *queue.Queue[*model.Task], store *store.Store, stop <-chan struct{}) {
	defer store.WG.Done()

	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d stopped work\n", workerID)
			return
		case task, exists := <-taskQueue.Dequeue():
			if !exists {
				return
			}
			store.Mu.Lock()
			task.Status = "IN_PROGRESS"
			store.Stats.InProgress++
			store.Mu.Unlock()

			fmt.Printf("Worker %d in progress on task %s\n", workerID, task.ID)

			time.Sleep(5 * time.Second)

			store.Mu.Lock()
			task.Status = "DONE"
			store.Stats.InProgress--
			store.Stats.Completed++
			store.Mu.Unlock()

			fmt.Printf("Worker %d finished task %s\n", workerID, task.ID)
		}
	}
}
