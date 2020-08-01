package types

import (
	"database/sql"
)

//Configuration Stores the main configuration for the application
type Configuration struct {
	ServerPort string
	SqlitePath string
	CsrfSecure bool
	Template   string
	Theme      string // couleur de base
}

//About as
type About struct {
	Application string
	Note        string
	Author      string
	Version     string
	Date        string
	Git         string
}

// Session as
type Session struct {
	Username string
	LoggedIn bool
}

/*
Package types is used to store the context struct which
is passed while templates are executed.
*/

//Order mapper avec la table "orders"
type Order struct {
	ID        int             `db:"orders_id,key,auto"`
	Name      sql.NullString  `db:"orders_name"`
	PtfID     string          `db:"orders_ptf_id"`
	Status    string          `db:"orders_order"`
	Time      string          `db:"orders_time"`
	Quote     sql.NullFloat64 `db:"orders_quote"`
	Quantity  sql.NullInt64   `db:"orders_quantity"`
	Buy       sql.NullFloat64 `db:"orders_buy"`
	Sell      sql.NullFloat64 `db:"orders_sell"`
	CostPrice sql.NullFloat64 `db:"orders_cost_price"`
	Cost      sql.NullFloat64 `db:"orders_cost"`
	Debit     sql.NullFloat64 `db:"orders_debit"`
	Credit    sql.NullFloat64 `db:"orders_credit"`
	Gain      sql.NullFloat64 `db:"orders_gain"`
	GainP     sql.NullFloat64 `db:"orders_gainp"`
	SellTime  sql.NullString  `db:"orders_sell_time"`
	SellCost  sql.NullFloat64 `db:"orders_sell_cost"`
	SellGain  sql.NullFloat64 `db:"orders_sell_gain"`
	SellGainP sql.NullFloat64 `db:"orders_sell_gainp"`
	Rem       sql.NullString  `db:"orders_rem"`
}

// TableName orders
func (*Order) TableName() string {
	return "orders"
}

// User compte utilisateur
type User struct {
	ID       int    `db:"user_id,key,auto"`
	Name     string `db:"user_name"`
	Password string `db:"user_password"`
	Email    string `db:"user_email"`
}

// TableName users
func (*User) TableName() string {
	return "users"
}

// //Context is the struct passed to templates
// type Context struct {
// 	Orders     []Order
// 	Navigation string
// 	Search     string
// 	Message    string
// 	CSRFToken  template.HTML
// 	Referer    string
// }

//Status is the JSON struct to be returned
type Status struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
