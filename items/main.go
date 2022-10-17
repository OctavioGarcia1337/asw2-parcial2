package main

import (
	"items/app"
	"items/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
