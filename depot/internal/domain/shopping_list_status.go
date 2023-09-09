package domain

type ShoppingListStatus string

const (
	ShoppingListUnknown     ShoppingListStatus = ""
	ShoppingListIsAvailable ShoppingListStatus = "available"
	ShoppingListIsAssigned  ShoppingListStatus = "assigned"
	ShoppingListIsActive    ShoppingListStatus = "active"
	ShoppingListIsCompleted ShoppingListStatus = "completed"
	ShoppingListIsCanceled  ShoppingListStatus = "cancelled"
)

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListUnknown, ShoppingListIsAvailable, ShoppingListIsAssigned, ShoppingListIsActive, ShoppingListIsCompleted,
		ShoppingListIsCanceled:
		return string(s)
	default:
		return ""
	}
}

func ToShoppingListStatus(status string) ShoppingListStatus {
	switch status {
	case ShoppingListIsAvailable.String():
		return ShoppingListIsAvailable
	case ShoppingListIsAssigned.String():
		return ShoppingListIsAssigned
	case ShoppingListIsAvailable.String():
		return ShoppingListIsAvailable
	case ShoppingListIsActive.String():
		return ShoppingListIsActive
	case ShoppingListIsCompleted.String():
		return ShoppingListIsCompleted
	case ShoppingListIsCanceled.String():
		return ShoppingListIsCanceled
	default:
		return ShoppingListUnknown
	}
}
