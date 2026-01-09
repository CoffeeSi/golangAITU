package model

type IDResponse struct {
	ID string `json:"id"`
}

type PostData struct {
	Payload string `json:"payload"`
}
type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Stats struct {
	Submitted  int `json:"submitted"`
	Completed  int `json:"completed"`
	InProgress int `json:"in_progress"`
}
