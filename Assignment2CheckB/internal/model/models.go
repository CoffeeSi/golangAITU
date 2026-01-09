package model

type Event struct {
	Type string
}

type Statistics struct {
	Received    uint64 `json:"received"`
	Processed   uint64 `json:"processed"`
	Queued      uint64 `json:"queued"`
	UniqueTypes uint64 `json:"unique_types"`
}
