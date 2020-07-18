package order

import (
    "errors"
)

func (order *OrderAggregate) PlaceOrder(cmd *PlaceOrderCommand) error {
    if order.OrderStatus != "" {
        return errors.New("Wrong status.")
    }
    if err := order.record(OrderPlacedEvent{OrderNumber: cmd.OrderNumber, ItemNo: cmd.OrderNumber}); err != nil {
        return err
    }
    return nil
}
