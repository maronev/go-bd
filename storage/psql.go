package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXIST products(
		id SERIAL NOT NULL,
		name VARCHAR2(25) NOT NULL,
		observations VARCHAR2(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
		`
)

// Usado para trabajar con postgres y el paquete product
type PsqlProduct struct {
	db *sql.DB
}

// Retorna un nuevo puntero de plsqlproduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Implementa la interfaz
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlCreateProduct)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	fmt.Println()

	return err
}
