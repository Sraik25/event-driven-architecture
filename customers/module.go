package customers

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/customers/internal/application"
	"github.com/Sraik25/event-driven-architecture/customers/internal/grpc"
	"github.com/Sraik25/event-driven-architecture/customers/internal/logging"
	"github.com/Sraik25/event-driven-architecture/customers/internal/postgres"
	"github.com/Sraik25/event-driven-architecture/customers/internal/rest"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/Sraik25/event-driven-architecture/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := postgres.NewCustomerRepository("customers.customers", mono.DB())

	// setup application
	var app application.App
	app = application.New(customers, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
