package es

import (
	"fmt"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
)

type (
	EventApplier interface {
		ApplyEvent(event ddd.Event) error
	}

	EventCommitter interface {
		CommitEvents()
	}
)

func LoadEvent(v interface{}, event ddd.AggregateEvent) error {
	type loader interface {
		EventApplier
		VersionSetter
	}

	agg, ok := v.(loader)

	if !ok {
		return fmt.Errorf("%T does not hav the methods implemented to load events", v)
	}

	if err := agg.ApplyEvent(event); err != nil {
		return err
	}

	agg.setVersion(event.AggregateVersion())

	return nil
}
