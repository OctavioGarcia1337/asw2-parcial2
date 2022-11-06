package main

import (
	"items/app"
	"items/db"
)

func main() {
	err := db.Init_db()
	if err != nil {
		return
	}
	app.StartRoute()
}
