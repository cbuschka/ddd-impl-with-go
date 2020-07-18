package order

import "fmt"

type InMemoryRepository struct {
    storedOrderEvents map[string]*[]OrderEvent
}

var repo = newRepository()

func newRepository() *InMemoryRepository {
    repo := InMemoryRepository{}
    repo.storedOrderEvents = make(map[string]*[]OrderEvent)
    return &repo
}

func GetInMemoryRepository() (Repository, error) {
    return repo, nil
}

func (repo *InMemoryRepository) FindOrder(orderNumber string) (*OrderAggregate, error) {
    events := repo.storedOrderEvents[orderNumber]
    if events == nil {
        return nil, fmt.Errorf("No such order: %s", orderNumber)
    }
    order := OrderAggregate{}
    if err := order.replay(events); err != nil {
        return nil, err
    }

    return &order, nil
}

func (repo *InMemoryRepository) StoreOrder(order *OrderAggregate) error {
    repo.storedOrderEvents[order.OrderNumber] = &order.events

    return nil
}
