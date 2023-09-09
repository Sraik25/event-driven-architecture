package stores

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/ddd"
	"github.com/Sraik25/event-driven-architecture/stores/internal/grpc"
	"github.com/Sraik25/event-driven-architecture/stores/internal/rest"

	"github.com/Sraik25/event-driven-architecture/internal/monolith"
	"github.com/Sraik25/event-driven-architecture/stores/internal/application"
	"github.com/Sraik25/event-driven-architecture/stores/internal/logging"
	"github.com/Sraik25/event-driven-architecture/stores/internal/postgres"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	stores := postgres.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := postgres.NewProductRepository("stores.product", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
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
