package order_test

import (
	"fmt"
	"testing"

	"github.com/DavesBorges/trade-platform/svc/order"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASellOrder(t *testing.T) {
	assert := assert.New(t)
	orderCreated, err := order.New("sell", "entity-1234", 13.4, 10)
	assert.Nil(err)
	assert.NotNil(orderCreated)

}

func TestShouldFailToCreateOrderWithInvalidInputs(t *testing.T) {
	type TestCase struct {
		kind      string
		entityId  string
		unitPrice float64
		quantity  int
	}

	var testCases []TestCase = []TestCase{
		{
			kind:      "sells",
			entityId:  "entity-1234",
			unitPrice: 10.2,
			quantity:  12,
		},
		{
			kind:      "sell",
			entityId:  "",
			unitPrice: 10.2,
			quantity:  12,
		},
		{
			kind:      "sell",
			entityId:  "entity-1234",
			unitPrice: 0,
			quantity:  12,
		},
		{
			kind:      "sell",
			entityId:  "entity-1234",
			unitPrice: 10,
			quantity:  0,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
			assert := assert.New(t)
			orderCreated, err := order.New(testCase.kind, testCase.entityId, testCase.unitPrice, testCase.quantity)
			assert.NotNil(err)
			assert.Nil(orderCreated)
		})

	}
}

func TestShouldCreateOrderWithValidInputs(t *testing.T) {
	type TestCase struct {
		kind      string
		entityId  string
		unitPrice float64
		quantity  int
	}

	var testCases []TestCase = []TestCase{
		{
			kind:      "sell",
			entityId:  "entity-1234",
			unitPrice: 1,
			quantity:  1,
		},
		{
			kind:      "sell",
			entityId:  "entity-1234",
			unitPrice: 100,
			quantity:  1,
		},
		{
			kind:      "sell",
			entityId:  "entity-1234",
			unitPrice: 100,
			quantity:  1,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
			assert := assert.New(t)
			orderCreated, err := order.New(testCase.kind, testCase.entityId, testCase.unitPrice, testCase.quantity)
			assert.Nil(err)
			assert.NotNil(orderCreated)
		})

	}
}
