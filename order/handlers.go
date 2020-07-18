package order

import (
    "errors"
    "fmt"
)

func (order *OrderAggregate) PlaceOrder(cmd *PlaceOrderCommand) error {
    if order.OrderStatus != "" {
        return errors.New("Wrong status.")
    }
    if err := order.add(OrderPlacedEvent{OrderNumber: cmd.OrderNumber, ItemNo: cmd.OrderNumber}); err != nil {
        return err
    }
    return nil
}

func (order *OrderAggregate) CancelOrder(cmd *CancelOrderCommand) error {
    if order.OrderStatus != OrderPlaced {
        return errors.New("Wrong status.")
    }
    if order.OrderNumber != cmd.OrderNumber {
        return fmt.Errorf("Invalid orderNumber: Expected %s, but was %s", order.OrderNumber, cmd.OrderNumber)
    }
    if err := order.add(OrderCancelledEvent{}); err != nil {
        return err
    }
    return nil
}
