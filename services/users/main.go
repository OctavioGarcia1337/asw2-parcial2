package main

import (
	"users/app"
	"users/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
