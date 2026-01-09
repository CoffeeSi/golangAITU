package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CoffeeSi/golangAITU/assignment2/internal/api"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/model"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/queue"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/store"
	"github.com/CoffeeSi/golangAITU/assignment2/internal/worker"
)

func main() {
	s := store.New()
	q := queue.New[*model.Task](100)

	stopWorkers := make(chan struct{})
	stopMonitor := make(chan struct{})

	for i := 1; i <= 2; i++ {
		s.WG.Add(1)
		go worker.Worker(i, q, s, stopWorkers)
	}

	go worker.Monitor(s, stopMonitor)

	handler := api.Handler{
		Store: s,
		Queue: q,
		Stop:  stopWorkers,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handler.TasksHandler)
	mux.HandleFunc("/tasks/{id}", handler.TaskHandler)
	mux.HandleFunc("/stats", handler.StatsHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
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
	close(stopMonitor)
	q.Close()

	fmt.Println("waiting for workers")
	s.WG.Wait()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)
	fmt.Println("exit")
}
