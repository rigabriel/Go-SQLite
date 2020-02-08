package debts

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Debts struct {
	DB *sql.DB
}

func (d *Debts) Printer(w http.ResponseWriter, r *http.Request) {
	jsonedDebt := d.Get()

	jsonDebt, err := json.Marshal(jsonedDebt)

	if err != nil {
		panic(err)
	}

	w.Write(jsonDebt)
}



func (d *Debts) Add(debt Debtor) {
	stmt, _ := d.DB.Prepare(`
		INSERT INTO debts (debtorName) values (?)
	`)

	stmt.Exec(debt.DebtorName)
}


func NewDebts(db *sql.DB) *Debts {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "debts" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"debtorName"	TEXT
		);
	`)

	stmt.Exec()

	return &Debts{
		DB: db,
	}
}

func (d *Debts) Get() []Debtor {
	var debtors []Debtor

	rows, _ := d.DB.Query(`
		SELECT * FROM debts
	`)

	var id int
	var debtorName string

	for rows.Next() {
		rows.Scan(&id, &debtorName)

		newDebtor := Debtor{
			ID:         id,
			DebtorName: debtorName,
		}
		debtors = append(debtors, newDebtor)
	}

	return debtors
}


