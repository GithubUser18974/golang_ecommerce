package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sikozonpc/ecom/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // import for rate limiting
	r.Use(middleware.RealIP)    // analytics and tracing
	r.Use(middleware.Logger)    // logs requests
	r.Use(middleware.Recoverer) // recovers from panic sverer

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All is Well :)"))
	})
	productService := products.NewService(app.db)
	producthandler := products.NewHandler(productService)

	r.Get("/products", producthandler.Listproducts)

	return r

}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server Has Started at Address %s", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config *config
	db     *sql.DB
}
type config struct {
	addr string
}
