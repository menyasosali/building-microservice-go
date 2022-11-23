package main

import (
	"building-microservices-go/product-api/handlers"
	"context"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := hclog.New(&hclog.LoggerOptions{
		Name:  "product-images",
		Level: hclog.LevelFromString("DEBUG"),
	})

	// create a logger for the server from the default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	// create the handlers
	np := handlers.NewProducts(sl)

	// create a new serve mux and register the handlers
	sw := mux.NewRouter()

	// handlers for API
	getRouter := sw.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", np.ListAll)
	getRouter.HandleFunc("/products/:{id[0-9]+}", np.ListSingle)

	putRouter := sw.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", np.Update)
	putRouter.Use(np.MiddlewareProductValidation)

	postRouter := sw.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", np.Create)
	postRouter.Use(np.MiddlewareProductValidation)

	deleteRouter := sw.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", np.DeleteProduct)

	// handlers for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// create a new server
	s := &http.Server{
		Addr:         ":9001",           // configure the bind address
		Handler:      ch(sw),            // set the default handler
		ErrorLog:     sl,                // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	// start the server
	go func() {
		sl.Printf("Starting server on port %s", s.Addr)

		err := s.ListenAndServe()
		if err != nil {
			sl.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	sl.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
