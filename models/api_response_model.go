package models

type APIRespone[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PagedAPIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Page    *int   `json:"page"`
	MaxPage int    `json:"max_page"`
	Data    T      `json:"data"`
}
