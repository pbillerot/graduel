package db

/*
	Accès à la table ORDERS
*/

import (
	"log"

	"github.com/pbillerot/graduel/types"
)

//GetOrders retrieves all the orders depending on the
//status sell or buy
func GetOrders(status string) ([]types.Order, error) {
	log.Println("GetOrders for ", status)
	db := openDb()
	orders := make([]types.Order, 0, 0)
	err = db.Select(&orders).Where("orders_order = ?", status).Do()
	if err != nil {
		log.Fatal(err)
	}
	return orders, nil
}
