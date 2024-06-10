package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/krutzschluca/LeaShifts/internal/model"
)

type API struct {
	m *model.Model
}

func New(m *model.Model) *API {
	return &API{
		m: m,
	}
}

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// CORS middleware to add necessary headers
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT")
		next.ServeHTTP(w, r)
	})
}

func (a *API) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	e, err := a.m.GetEmployee(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get employee: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, e)
}

func (a *API) GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	employees, err := a.m.GetEmployees(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get employees: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, employees)
}

func (a *API) AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var e model.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "could not decode employee", http.StatusBadRequest)
		return
	}

	err = a.m.AddEmployee(r.Context(), &e)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not add employee: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, e)
}

func (a *API) UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var e model.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "could not decode employee", http.StatusBadRequest)
		return
	}
	stringID := r.PathValue("id")
	e.ID, err = strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.UpdateEmployee(r.Context(), &e)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not update employee: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, e)
}

func (a *API) DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.DeleteEmployee(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not delete employee: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *API) GetStoreHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	s, err := a.m.GetStore(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get store: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) GetStoresHandler(w http.ResponseWriter, r *http.Request) {
	stores, err := a.m.GetStores(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get stores: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, stores)
}

func (a *API) AddStoreHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Store
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode store: %v", err), http.StatusBadRequest)
		return
	}

	err = a.m.AddStore(r.Context(), &s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not add store: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) UpdateStoreHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Store
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode store: %v", err), http.StatusBadRequest)
		return
	}
	stringID := r.PathValue("id")
	s.ID, err = strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.UpdateStore(r.Context(), &s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not update store: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) DeleteStoreHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.DeleteStore(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not delete store: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *API) GetShiftHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	s, err := a.m.GetShift(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get shift: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) GetShiftsHandler(w http.ResponseWriter, r *http.Request) {
	shifts, err := a.m.GetShifts(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("could not get shifts: %v", err), http.StatusInternalServerError)
		return
	}

	re

func (a *API) AddShiftHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Shift
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode shift: %v", err), http.StatusBadRequest)
		return
	}

	err = a.m.AddShift(r.Context(), &s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not add shift: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) UpdateShiftHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Shift
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not decode shift: %v", err), http.StatusBadRequest)
		return
	}
	stringID := r.PathValue("id")
	s.ID, err = strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.UpdateShift(r.Context(), &s)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not update shift: %v", err), http.StatusInternalServerError)
		return
	}

	renderJSON(w, s)
}

func (a *API) DeleteShiftHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = a.m.DeleteShift(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not delete shift: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
