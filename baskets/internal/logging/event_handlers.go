package logging

import (
	"context"
	"github.com/rs/zerolog"

	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type EventHandler[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger zerolog.Logger
}

func LogEventHandlerAccess[T ddd.Event](eventHandler ddd.EventHandler[T], label string, logger zerolog.Logger) EventHandler[T] {
	return EventHandler[T]{
		EventHandler: eventHandler,
		label:        label,
		logger:       logger,
	}
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandler[ddd.Event])(nil)

func (h EventHandler[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Info().Msgf("--> Baskets.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- Baskets.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
