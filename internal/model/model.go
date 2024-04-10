package model

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type Model struct {
	dbConnection *sqlx.DB
}

type Store struct {
	ID          int    `json:"id"                     db:"id"`
	Address     string `json:"address"                db:"address"`
	OpeningTime string `json:"opening_time,omitempty" db:"opening_time"`
	ClosingTime string `json:"closing_time,omitempty" db:"closing_time"`
}

type Employee struct {
	ID   int    `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

type Shift struct {
	ID           int       `json:"id"            db:"id"`
	StoreID      int       `json:"store_id"      db:"store_id"`
	EmployeeID   int       `json:"employee_id"   db:"employee_id"`
	Date         time.Time `json:"date"          db:"date"`
	StartingTime string    `json:"starting_time" db:"starting_time"`
	EndingTime   string    `json:"ending_time"   db:"ending_time"`
}

func New(dbConnection *sqlx.DB) *Model {
	return &Model{
		dbConnection: dbConnection,
	}
}

func (a *Model) GetEmployee(ctx context.Context, id int) (*Employee, error) {
	var e Employee
	err := a.dbConnection.GetContext(ctx, &e, "SELECT * FROM employees WHERE id=$1", id)
	return &e, err
}

func (a *Model) GetEmployees(ctx context.Context) ([]Employee, error) {
	var employees []Employee
	err := a.dbConnection.SelectContext(ctx, &employees, "SELECT * FROM employees")
	return employees, err
}

func (a *Model) AddEmployee(ctx context.Context, e *Employee) error {
	err := a.dbConnection.GetContext(ctx, e, "INSERT INTO employees (name) VALUES ($1) RETURNING *", e.Name)
	return err
}

func (a *Model) UpdateEmployee(ctx context.Context, e *Employee) error {
	_, err := a.dbConnection.ExecContext(ctx, "UPDATE employees SET name=$1 WHERE id=$2", e.Name, e.ID)
	return err
}

func (a *Model) DeleteEmployee(ctx context.Context, id int) error {
	_, err := a.dbConnection.ExecContext(ctx, "DELETE FROM employees WHERE id=$1", id)
	return err
}

func (a *Model) GetStore(ctx context.Context, id int) (*Store, error) {
	var s Store
	err := a.dbConnection.GetContext(ctx, &s, "SELECT * FROM stores WHERE id=$1", id)
	return &s, err
}

func (a *Model) GetStores(ctx context.Context) ([]Store, error) {
	var stores []Store
	err := a.dbConnection.SelectContext(ctx, &stores, "SELECT * FROM stores")
	return stores, err
}

func (a *Model) AddStore(ctx context.Context, s *Store) error {
	err := a.dbConnection.GetContext(ctx, s, "INSERT INTO stores (address, opening_time, closing_time) VALUES ($1, $2, $3) RETURNING *", s.Address, s.OpeningTime, s.ClosingTime)
	return err
}

func (a *Model) UpdateStore(ctx context.Context, s *Store) error {
	_, err := a.dbConnection.ExecContext(ctx, "UPDATE stores SET address=$1, opening_time=$2, closing_time=$3 WHERE id=$4", s.Address, s.OpeningTime, s.ClosingTime, s.ID)
	return err
}

func (a *Model) DeleteStore(ctx context.Context, id int) error {
	_, err := a.dbConnection.ExecContext(ctx, "DELETE FROM stores WHERE id=$1", id)
	return err
}

func (a *Model) GetShift(ctx context.Context, id int) (*Shift, error) {
	var s Shift
	err := a.dbConnection.GetContext(ctx, &s, "SELECT * FROM shifts WHERE id=$1", id)
	return &s, err
}

func (a *Model) GetShifts(ctx context.Context) ([]Shift, error) {
	var shifts []Shift
	err := a.dbConnection.SelectContext(ctx, &shifts, "SELECT * FROM shifts")
	return shifts, err
}

func (a *Model) AddShift(ctx context.Context, s *Shift) error {
	err := a.dbConnection.GetContext(ctx, s, "INSERT INTO shifts (store_id, employee_id, date, starting_time, ending_time) VALUES ($1, $2, $3, $4, $5) RETURNING *", s.StoreID, s.EmployeeID, s.Date, s.StartingTime, s.EndingTime)
	return err
}

func (a *Model) UpdateShift(ctx context.Context, s *Shift) error {
	_, err := a.dbConnection.ExecContext(ctx, "UPDATE shifts SET store_id=$1, employee_id=$2, date=$3, starting_time=$4, ending_time=$5 WHERE id=$6", s.StoreID, s.EmployeeID, s.Date, s.StartingTime, s.EndingTime, s.ID)
	return err
}

func (a *Model) DeleteShift(ctx context.Context, id int) error {
	_, err := a.dbConnection.ExecContext(ctx, "DELETE FROM shifts WHERE id	=$1", id)
	return err
}
