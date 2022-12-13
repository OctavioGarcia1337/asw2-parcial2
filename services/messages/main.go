package main

import (
	"messages/app"
	"messages/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
