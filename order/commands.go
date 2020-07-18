package order

type PlaceOrderCommand struct {
    OrderNumber string
    ItemNo      string
}

type MarkOrderDeliveredCommand struct {
    OrderNumber string
}

type CancelOrderCommand struct {
    OrderNumber string
}
