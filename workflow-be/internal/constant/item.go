package constant

type ItemStatus string 
const (
	ItemApprovedStatus ItemStatus = "APPROVED"
	ItemRejectedStatus ItemStatus = "REJECTED"
	ItemPenddingStatus  ItemStatus = "PENDING"
)