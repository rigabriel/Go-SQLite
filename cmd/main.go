package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"../platform/debts"
)

func main() {


	db, _ := sql.Open("sqlite3", "./debts.db")

	debt := debts.NewDebts(db)

	debt.Add(debts.Debtor {
		DebtorName: "nicolas.walter",
	})

	debtors := debt.Get()

	fmt.Println(debtors)

	d := &debts.Debts{
		DB: db,
	}

	http.HandleFunc("/debtors", d.Printer)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

