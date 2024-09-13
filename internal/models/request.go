package models

type RequestItem struct {
	Title     string
	Amount    float64
	Quantity  uint
}


type RequestItemWithID struct {
	ID       uint    `json:"id" binding:"required"`
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Quantity uint    `json:"quantity"`
}

type RequestGetItems struct {
	Items []RequestItem `json:"items" binding:"required"`
}