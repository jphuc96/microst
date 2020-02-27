package models

type SystemStatus struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
