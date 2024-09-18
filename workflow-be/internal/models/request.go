package models

import "workflow/internal/constant"

type RequestItem struct {
	Title     string `json:"title"`
	Amount    float64 `json:"amount"`
	Quantity  uint `json:"quantity"`
}


type RequestItemWithID struct {
	ID       uint    `json:"id" binding:"required"`
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Quantity uint    `json:"quantity"`
}

type RequestUpdateItemStatus struct {
	Status constant.ItemStatus
}

type RequestGetItems struct {
	Items []RequestItem `json:"items" binding:"required"`
}
type RequestLogin struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}