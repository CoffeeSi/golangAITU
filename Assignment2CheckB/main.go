package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/api"
	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/store"
	"github.com/CoffeeSi/golangAITU/assignment2CheckB/internal/worker"
)

func main() {
	var WG sync.WaitGroup
	s := store.NewCounterStore[model.Event]()
	stats := &model.Statistics{
		Received:  0,
		Processed: 0,
	}

	stopWorkers := make(chan struct{})
	eventChan := make(chan model.Event, 100)

	for i := 1; i <= 2; i++ {
		w := &worker.Worker{
			ID:    i,
			Stats: stats,
			Store: s,
			WG:    &WG,
			Stop:  stopWorkers,
		}
		fmt.Printf("Worker %d started\n", i)
		WG.Add(1)
		go w.Start(eventChan)
	}

	handler := api.Handler{
		Store:     s,
		Stats:     stats,
		Stop:      stopWorkers,
		EventChan: eventChan,
	}
	http.HandleFunc("/events", handler.EventsHandler)
	http.HandleFunc("/counts", handler.CountsHandler)
	http.HandleFunc("/stats", handler.StatsHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		fmt.Println("http://localhost:8080")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("error: ", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs

	close(stopWorkers)

	fmt.Println("Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}

	close(eventChan)

	fmt.Println("Waiting for workers to finish...")
	WG.Wait()

	fmt.Println("Shutdown complete")
}
