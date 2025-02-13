package repository

import (
	"database/sql"
	"farmish/model"

	"github.com/google/uuid"
)

type Warehouse interface {
	InsertStock(model.CreateStockDTO) error
	FindStockByName(string) (model.Stock, error)
	UpdateFeedStock(model.SupplyStockDTO) error
	UpdateMedicineStock(model.SupplyStockDTO) error
	FindFeed() ([]model.Stock, error)
	FindMedicine() ([]model.Stock, error)
}
type WarehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db}
}

func (w *WarehouseRepository) FindFeed() ([]model.Stock, error) {
	var allFeed []model.Stock = []model.Stock{}
	rows, err := w.db.Query(
		`SELECT id, name, quantity, cost, stock_type
		FROM warehouse WHERE stock_type = $1`,
		"feed",
	)
	if err != nil {
		return allFeed, err
	}
	defer rows.Close()
	for rows.Next() {
		var feed model.Stock
		if err := rows.Scan(
			&feed.Id,
			&feed.Name,
			&feed.Quantity,
			&feed.Cost,
			&feed.StockType,
		); err != nil {
			return allFeed, err
		}
		allFeed = append(allFeed, feed)
	}

	return allFeed, nil
}

func (w *WarehouseRepository) FindMedicine() ([]model.Stock, error) {
	var allMedicine []model.Stock = []model.Stock{}
	rows, err := w.db.Query(
		`SELECT id, name, quantity, cost, stock_type
		FROM warehouse WHERE stock_type = $1`,
		"medicine",
	)
	if err != nil {
		return allMedicine, err
	}
	defer rows.Close()
	for rows.Next() {
		var medicine model.Stock
		if err := rows.Scan(
			&medicine.Id,
			&medicine.Name,
			&medicine.Quantity,
			&medicine.Cost,
			&medicine.StockType,
		); err != nil {
			return allMedicine, err
		}
		allMedicine = append(allMedicine, medicine)
	}

	return allMedicine, nil
}

func (w *WarehouseRepository) InsertStock(newStock model.CreateStockDTO) error {
	_, err := w.db.Exec(`
	INSERT INTO warehouse (id, name, quantity, cost, stock_type) 
	VALUES ($1, $2, $3, $4, $5)
	`,
		uuid.NewString(),
		newStock.Name,
		newStock.Quantity,
		newStock.Cost,
		newStock.StockType,
	)
	return err
}

func (w *WarehouseRepository) FindStockByName(name string) (model.Stock, error) {
	var stock model.Stock

	row := w.db.QueryRow(
		`SELECT id, name, quantity, cost, stockType FROM warehouse WHERE name = $1`,
		name,
	)

	err := row.Scan(
		&stock.Id,
		&stock.Name,
		&stock.Quantity,
		&stock.Cost,
		&stock.StockType,
	)
	return stock, err
}

func (w *WarehouseRepository) UpdateFeedStock(stock model.SupplyStockDTO) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(
		`INSERT INTO feed_arrival VALUES ($1, $2, $3, $4, $5)`,
		uuid.NewString(),
		stock.Type,
		stock.Quantity,
		stock.ArrivalDate,
		stock.Notes,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE warehouse SET quantity = quantity + $1 WHERE name = $2`,
		stock.Quantity,
		stock.Type,
	)

	if err != nil {
		return err
	}

	return nil
}

func (w *WarehouseRepository) UpdateMedicineStock(stock model.SupplyStockDTO) error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(
		`INSERT INTO medicine_arrival VALUES ($1, $2, $3, $4, $5)`,
		uuid.NewString(),
		stock.Type,
		stock.Quantity,
		stock.ArrivalDate,
		stock.Notes,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE warehouse SET quantity = quantity + $1 WHERE name = $2`,
		stock.Quantity,
		stock.Type,
	)

	if err != nil {
		return err
	}

	return nil
}
