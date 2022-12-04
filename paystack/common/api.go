package common

type ApiResponse[T any] struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}
