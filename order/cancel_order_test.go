package order

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCanCancelOrderWhenPlaced(t *testing.T) {
    order := NewOrder()
    order.OrderStatus = OrderPlaced
    order.OrderNumber = "O1"

    cmd := CancelOrderCommand{OrderNumber: order.OrderNumber}
    err := order.CancelOrder(&cmd)

    assert.Nil(t, err)
    assert.Equal(t, cmd.OrderNumber, order.OrderNumber)
    assert.Equal(t, OrderCancelled, order.OrderStatus)
}

func TestCannotCancelOrderWhenCancelled(t *testing.T) {
    order := NewOrder()
    order.OrderStatus = OrderCancelled
    order.OrderNumber = "O1"

    cmd := CancelOrderCommand{OrderNumber: order.OrderNumber}
    err := order.CancelOrder(&cmd)

    assert.NotNil(t, err)
}

func TestCannotCancelOrderWhenNew(t *testing.T) {
    order := NewOrder()
    order.OrderStatus = ""
    order.OrderNumber = "O1"

    cmd := CancelOrderCommand{OrderNumber: order.OrderNumber}
    err := order.CancelOrder(&cmd)

    assert.NotNil(t, err)
}
