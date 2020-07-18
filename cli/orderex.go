package main

import (
    o "github.com/cbuschka/ddd-in-go/order"
)

func placeOrder() (string, error) {
    order := o.NewOrder()
    if err := order.PlaceOrder(&o.PlaceOrderCommand{OrderNumber: "O1"}); err != nil {
        return "", err
    }
    if err := o.StoreOrder(order); err != nil {
        return "", err
    }
    return order.OrderNumber, nil
}

func cancelOrder(orderNumber string) error {
    order, err := o.FindOrder(orderNumber)
    if err != nil {
        panic(err.Error())
    }
    if err := order.CancelOrder(&o.CancelOrderCommand{OrderNumber: orderNumber}); err != nil {
        return err
    }
    if err := o.StoreOrder(order); err != nil {
        return err
    }

    return nil
}

func main() {
    orderNumber, err := placeOrder()
    if err != nil {
        panic(err.Error())
    }

    if err := cancelOrder(orderNumber); err != nil {
        panic(err.Error())
    }
}
