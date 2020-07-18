package order

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCanPlaceOrderWhenNew(t *testing.T) {
    order := NewOrder()
    assert.Equal(t, OrderStatus(""), order.OrderStatus)

    cmd := PlaceOrderCommand{OrderNumber: "O2"}
    err := order.PlaceOrder(&cmd)

    assert.Nil(t, err)
    assert.Equal(t, cmd.OrderNumber, order.OrderNumber)
}

func TestCannotPlaceOrderWhenPlaced(t *testing.T) {
    order := NewOrder()
    order.OrderStatus = OrderPlaced
    order.OrderNumber = "O1"

    cmd := PlaceOrderCommand{OrderNumber: "O1"}
    err := order.PlaceOrder(&cmd)

    assert.NotNil(t, err)
}
