package main

import (
	"database/sql"
	"fmt"
	"github.com/Sraik25/event-driven-architecture/internal/web"
	"google.golang.org/grpc/reflection"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"

	"github.com/Sraik25/event-driven-architecture/baskets"
	"github.com/Sraik25/event-driven-architecture/customers"
	"github.com/Sraik25/event-driven-architecture/depot"
	"github.com/Sraik25/event-driven-architecture/internal/config"
	"github.com/Sraik25/event-driven-architecture/internal/logger"
	"github.com/Sraik25/event-driven-architecture/internal/monolith"
	"github.com/Sraik25/event-driven-architecture/internal/rpc"
	"github.com/Sraik25/event-driven-architecture/internal/waiter"
	"github.com/Sraik25/event-driven-architecture/notifications"
	"github.com/Sraik25/event-driven-architecture/ordering"
	"github.com/Sraik25/event-driven-architecture/payments"
	"github.com/Sraik25/event-driven-architecture/stores"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig

	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	m.db, err = sql.Open("pgx", cfg.PG.Coon)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(m.db)

	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})

	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Rpc)
	m.waiter = waiter.New(waiter.CatchSignal())

	m.modules = []monolith.Module{
		&baskets.Module{},
		&customers.Module{},
		&depot.Module{},
		&notifications.Module{},
		&ordering.Module{},
		&payments.Module{},
		&stores.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbts application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initRpc(rpc rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)
	return server
}

func initMux(rpcConfig rpc.RpcConfig) *chi.Mux {
	return chi.NewMux()
}
