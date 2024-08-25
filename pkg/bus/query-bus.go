package bus

import (
	"fmt"
)

type QueryBus struct {
	handlers map[string]QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (b *QueryBus) Register(query Query, handler QueryHandler) {
	name := fmt.Sprintf("%T", query)
	b.handlers[name] = handler
}

// Dispatch a query to its respective handler.
func (b *QueryBus) Dispatch(query Query) (interface{}, error) {
	name := fmt.Sprintf("%T", query)
	handler, ok := b.handlers[name]
	if !ok {
		return nil, fmt.Errorf("no handlers registered for query of type %s", name)
	}

	return handler.Handle(query)
}

// Query interface
type Query interface {
	Data() map[string]interface{}
}

// QueryHandler interface
type QueryHandler interface {
	Handle(query Query) (interface{}, error)
}
