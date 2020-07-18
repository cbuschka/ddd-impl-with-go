package order

type OrderEvent interface {
}

type OrderPlacedEvent struct {
    OrderNumber string
    ItemNo      string
}

type OrderCancelledEvent struct {
}
