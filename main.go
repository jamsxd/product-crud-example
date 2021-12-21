package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/log"
	"github.com/jamsxd/product-crud-example/pkg/product/application"
	"github.com/jamsxd/product-crud-example/pkg/product/domain"
	"github.com/jamsxd/product-crud-example/pkg/product/infrastructure/persistence"
	"github.com/jamsxd/product-crud-example/pkg/product/infrastructure/transport"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		port         = os.Getenv("PORT")
		dbName       = os.Getenv("DB_NAME")
		dbCollection = os.Getenv("DB_COLLECTION")
		dbHost       = os.Getenv("DB_HOST")
		dbUser       = os.Getenv("DB_USER")
		dbPass       = os.Getenv("DB_PASS")
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conexion := "mongodb://" + dbUser + ":" + dbPass + "@" + dbHost + "/"
	fmt.Println(conexion)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conexion))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	var (
		repo          = persistence.NewMongoRepository(client, dbName, dbCollection)
		svc           = domain.NewProductService(repo, logger)
		endpoint      = application.NewProductEndpoint(svc)
		httpTransport = transport.NewHttpHandler(endpoint, logger)
	)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		_ = logger.Log("transport", "HTTP", "addr", ":"+port)
		errs <- http.ListenAndServe("0.0.0.0:"+port, httpTransport)
	}()

	_ = logger.Log("exit", <-errs)

}
