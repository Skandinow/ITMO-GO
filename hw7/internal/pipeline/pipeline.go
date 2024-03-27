package pipeline

import (
	"context"

	"hw7/internal/model"
)

const fanLimit = 10

type OrderPipeline interface {
	Start(ctx context.Context, actions model.OrderActions, orders <-chan model.OrderInitialized, processed chan<- model.OrderProcessFinished)
}

type OrderPipelineImplementation struct{}

func NewOrderPipelineImplementation() *OrderPipelineImplementation {
	return &OrderPipelineImplementation{}
}

func (o *OrderPipelineImplementation) Start(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
	processed chan<- model.OrderProcessFinished,
) {
}
