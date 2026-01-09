package worker

import (
	"fmt"
	"time"

	"github.com/CoffeeSi/golangAITU/assignment2/internal/store"
)

func Monitor(store *store.Store, stop <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			fmt.Printf("Monitoring stopped\n")
			return
		case <-ticker.C:
			store.Mu.Lock()
			submitted := store.Stats.Submitted
			inProgress := store.Stats.InProgress
			completed := store.Stats.InProgress
			store.Mu.Unlock()

			pending := submitted - (inProgress + completed)
			fmt.Printf("PENDING: %d IN PROGRESS: %d DONE: %d\n", pending, inProgress, completed)
		}
	}
}
