package database

import (
	"database/sql"
	"fmt"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"strconv"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll() []entity.Order {
	var orders []entity.Order
	rows, err := r.Db.Query(`Select * from orders`)
	if err != nil {
		panic(err)
	}

	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			panic(err)
		}
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
		price, _ := strconv.ParseFloat(result[1], 64)
		tax, _ := strconv.ParseFloat(result[2], 64)
		finalPrice, _ := strconv.ParseFloat(result[3], 64)
		orders = append(orders, entity.Order{
			ID:         result[0],
			Price:      price,
			Tax:        tax,
			FinalPrice: finalPrice,
		})
	}

	return orders
}
