package order

import (
	"errors"
	"fmt"
)

type Order struct {
	Kind      string  `json:"kind"`
	EntityId  string  `json:"entityId"`
	UnitPrice float64 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
}

func New(kind string, entityId string, unitPrice float64, quantity int) (*Order, error) {
	if kind != "sell" {
		return nil, errors.New(fmt.Sprintf("Order kind must be of type 'sell'. %v was provided", kind))
	}

	if entityId == "" {
		return nil, errors.New("Invalid entity id \"\"")
	}

	if unitPrice <= 0 {
		return nil, errors.New(fmt.Sprintf("Unit price must be positive. %v was provided", unitPrice))
	}

	if quantity <= 0 {
		return nil, errors.New(fmt.Sprintf("Quantity must be positive. %v was provided", unitPrice))
	}
	return &Order{Kind: kind, EntityId: entityId, UnitPrice: unitPrice, Quantity: quantity}, nil
}
