package main

import (
    o "github.com/cbuschka/ddd-in-go/order"
)

func placeOrder(repo o.Repository) (string, error) {
    order := o.NewOrder()
    if err := order.PlaceOrder(&o.PlaceOrderCommand{OrderNumber: "O1"}); err != nil {
        return "", err
    }
    if err := repo.StoreOrder(order); err != nil {
        return "", err
    }
    return order.OrderNumber, nil
}

func cancelOrder(orderNumber string, repo o.Repository) error {
    order, err := repo.FindOrder(orderNumber)
    if err != nil {
        panic(err.Error())
    }
    if err := order.CancelOrder(&o.CancelOrderCommand{OrderNumber: orderNumber}); err != nil {
        return err
    }
    if err := repo.StoreOrder(order); err != nil {
        return err
    }

    return nil
}

func main() {
    repo, err := o.GetInMemoryRepository()
    if err != nil {
        panic(err.Error())
    }
    orderNumber, err := placeOrder(repo)
    if err != nil {
        panic(err.Error())
    }

    if err := cancelOrder(orderNumber, repo); err != nil {
        panic(err.Error())
    }
}
