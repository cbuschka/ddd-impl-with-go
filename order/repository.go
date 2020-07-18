package order

type Repository interface {
    FindOrder(orderNumber string) (*OrderAggregate, error)
    StoreOrder(order *OrderAggregate) error
}
