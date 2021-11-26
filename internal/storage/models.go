package storage

type RequestStatus struct {
	RequestID string `json:"request_id" db:"request_id"`
	Status    string `json:"status" db:"status"`
}
