package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	externaldatasource "github.com/fernandormoraes/go-clean-architecture/data/datasources/external"
	datarepository "github.com/fernandormoraes/go-clean-architecture/data/repositories"
	"github.com/fernandormoraes/go-clean-architecture/domain/usecases"

	bmhttp "github.com/fernandormoraes/go-clean-architecture/bookmark/delivery/http"
)

type App struct {
	httpServer *http.Server

	bookmarkUC usecases.BookmarkUseCase
}

func NewApp() *App {
	db := initDB()

	bookmarkExternalDatasource := externaldatasource.NewDbBookmarkDatasource(db, viper.GetString("mongo.bookmark_collection"))
	bookmarkRepo := datarepository.NewBookmarkRepository(bookmarkExternalDatasource)

	return &App{
		bookmarkUC: *usecases.NewBookmarkUseCase(bookmarkRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Database connected")

	return client.Database(viper.GetString("mongo.name"))
}
