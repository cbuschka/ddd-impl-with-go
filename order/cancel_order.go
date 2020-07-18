package order

import (
    "errors"
    "fmt"
)

func (order *OrderAggregate) CancelOrder(cmd *CancelOrderCommand) error {
    if order.OrderStatus != OrderPlaced {
        return errors.New("Wrong status.")
    }
    if order.OrderNumber != cmd.OrderNumber {
        return fmt.Errorf("Invalid orderNumber: Expected %s, but was %s", order.OrderNumber, cmd.OrderNumber)
    }
    if err := order.record(OrderCancelledEvent{}); err != nil {
        return err
    }
    return nil
}
