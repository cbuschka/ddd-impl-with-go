package order

type PlaceOrderCommand struct {
    OrderNumber string
    ItemNo      string
}

type CancelOrderCommand struct {
    OrderNumber string
}
