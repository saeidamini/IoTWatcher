package models

type DeviceStateLog struct {
	DeviceID    string `json:"DeviceID"`
	StateDate   string `json:"State#Date"`
	Operator    string `json:"Operator"`
	Date        string `json:"Date"`
	State       string `json:"State"`
	EscalatedTo string `json:"EscalatedTo"`
}
