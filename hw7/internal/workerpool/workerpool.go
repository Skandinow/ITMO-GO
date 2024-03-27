package workerpool

import (
	"context"

	"hw7/internal/model"
)

type OrderWorkerPool interface {
	StartWorkerPool(ctx context.Context, orders <-chan model.OrderInitialized, additionalActions model.OrderActions, workersCount int) <-chan model.OrderProcessFinished
}

type OrderWorkerPoolImplementation struct{}

func NewOrderWorkerPoolImplementation() *OrderWorkerPoolImplementation {
	return &OrderWorkerPoolImplementation{}
}

func (o *OrderWorkerPoolImplementation) StartWorkerPool(ctx context.Context, orders <-chan model.OrderInitialized, additionalActions model.OrderActions, workersCount int) <-chan model.OrderProcessFinished {
}
