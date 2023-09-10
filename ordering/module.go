package ordering

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/internal/monolith"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/application"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/grpc"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/logging"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/postgres"
	"github.com/Sraik25/event-driven-architecture/ordering/internal/rest"
)

type Module struct {
}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	orders := postgres.NewOrderRepository("ordering.orders", mono.DB())
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, invoices, shopping, notifications)
	app = logging.NewApplication(app, mono.Logger())

	// setup Driver adapters
	if err = grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}

	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}

	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
