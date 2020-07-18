package order

import (
    "fmt"
    "reflect"
)

type OrderStatus string

const (
    OrderPlaced    OrderStatus = "OrderPlaced"
    OrderCancelled OrderStatus = "OrderCancelled"
)

type OrderAggregate struct {
    OrderNumber string
    ItemNumber  string
    OrderStatus OrderStatus
    events      []OrderEvent
}

func NewOrder() *OrderAggregate {
    return &OrderAggregate{}
}

func (order *OrderAggregate) apply(event OrderEvent) error {
    switch event.(type) {
    case OrderPlacedEvent:
        order.OrderNumber = event.(OrderPlacedEvent).OrderNumber
        order.OrderStatus = OrderPlaced
        return nil
    case OrderCancelledEvent:
        order.OrderStatus = OrderCancelled
        return nil
    default:
        return fmt.Errorf("Unkown event type: %s", reflect.TypeOf(event))
    }
}

func (order *OrderAggregate) add(event OrderEvent) error {
    if err := order.apply(event); err != nil {
        return err
    }
    order.events = append(order.events, event)
    return nil
}

func (order *OrderAggregate) replay(events *[]OrderEvent) error {
    for _, event := range *events {
        if err := order.apply(event); err != nil {
            return err
        }
    }
    order.events = *events
    return nil
}
