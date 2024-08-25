package bus

import "fmt"

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{make(map[string]CommandHandler)}
}

func (b *CommandBus) Register(command Command, handler CommandHandler) {
	name := fmt.Sprintf("%T", command)
	b.handlers[name] = handler
}

func (b *CommandBus) Send(command Command) error {
	name := fmt.Sprintf("%T", command)
	if h, ok := b.handlers[name]; ok {
		return h.Handle(command)
	}
	return fmt.Errorf("no handler found for command %s", name)
}

type Command interface {
	ID() string
	Data() map[string]interface{}
}

type CommandHandler interface {
	Handle(command Command) error
}
