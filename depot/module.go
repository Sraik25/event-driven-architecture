package depot

import (
	"context"

	"github.com/Sraik25/event-driven-architecture/depot/internal/application"
	"github.com/Sraik25/event-driven-architecture/depot/internal/grpc"
	"github.com/Sraik25/event-driven-architecture/depot/internal/handlers"
	"github.com/Sraik25/event-driven-architecture/depot/internal/logging"
	"github.com/Sraik25/event-driven-architecture/depot/internal/postgres"
	"github.com/Sraik25/event-driven-architecture/depot/internal/rest"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/Sraik25/event-driven-architecture/internal/monolith"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	shoppingLists := postgres.NewShoppingListRepository("depot.shopping_lists", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	stores := grpc.NewStoreRepository(conn)
	products := grpc.NewProductRepository(conn)
	orders := grpc.NewOrderRepository(conn)

	// setup application
	var app application.App
	app = application.New(shoppingLists, stores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orders),
		mono.Logger(),
	)

	// setup Driver adapters
	if err = grpc.Register(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)

	return nil
}
