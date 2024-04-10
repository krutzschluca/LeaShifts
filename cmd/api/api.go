package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	config "github.com/luca/webdev/cmd/api/internal"
	"github.com/luca/webdev/internal/api"
	"github.com/luca/webdev/internal/model"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // nolint:errcheck
	logger.Info("starting")

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("could not create a config", zap.Error(err))
	}

	dbConnection, err := sqlx.Connect(cfg.DBDriver, cfg.DBConnectionString)
	if err != nil {
		logger.Fatal("could not create a database connection", zap.Error(err))
	}

	m := model.New(dbConnection)
	a := api.New(m)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.GetEmployeeHandler(w, r)
	})
	mux.HandleFunc("POST /employees", func(w http.ResponseWriter, r *http.Request) {
		a.AddEmployeeHandler(w, r)
	})
	mux.HandleFunc("GET /employees", func(w http.ResponseWriter, r *http.Request) {
		a.GetEmployeesHandler(w, r)
	})
	mux.HandleFunc("PUT /employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.UpdateEmployeeHandler(w, r)
	})
	mux.HandleFunc("DELETE /employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.DeleteEmployeeHandler(w, r)
	})

	mux.HandleFunc("GET /stores/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.GetStoreHandler(w, r)
	})
	mux.HandleFunc("GET /stores", func(w http.ResponseWriter, r *http.Request) {
		a.GetStoresHandler(w, r)
	})
	mux.HandleFunc("POST /stores", func(w http.ResponseWriter, r *http.Request) {
		a.AddStoreHandler(w, r)
	})
	mux.HandleFunc("PUT /stores/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.UpdateStoreHandler(w, r)
	})
	mux.HandleFunc("DELETE /stores/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.DeleteStoreHandler(w, r)
	})

	mux.HandleFunc("GET /shifts/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.GetShiftHandler(w, r)
	})
	mux.HandleFunc("GET /shifts", func(w http.ResponseWriter, r *http.Request) {
		a.GetShiftsHandler(w, r)
	})
	mux.HandleFunc("POST /shifts", func(w http.ResponseWriter, r *http.Request) {
		a.AddShiftHandler(w, r)
	})
	mux.HandleFunc("PUT /shifts/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.UpdateShiftHandler(w, r)
	})
	mux.HandleFunc("DELETE /shifts/{id}", func(w http.ResponseWriter, r *http.Request) {
		a.DeleteShiftHandler(w, r)
	})

	http.ListenAndServe("localhost"+cfg.Listen, mux) // nolint:errcheck
}
