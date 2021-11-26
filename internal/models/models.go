package models

type Scenario struct {
	ID         int64  `json:"id" db:"id" `
	Name       string `json:"name" db:"id"`
	ScenarioID string `json:"scenario_id" db:"id"`
}

