package order

import "fmt"

var storedOrderEvents = make(map[string]*[]OrderEvent)

func FindOrder(orderNumber string) (*OrderAggregate, error) {
    events := storedOrderEvents[orderNumber]
    if events == nil {
        return nil, fmt.Errorf("No such order: %s", orderNumber)
    }
    order := OrderAggregate{}
    if err := order.replay(events); err != nil {
        return nil, err
    }

    return &order, nil
}

func StoreOrder(order *OrderAggregate) error {
    storedOrderEvents[order.OrderNumber] = &order.events

    return nil
}
