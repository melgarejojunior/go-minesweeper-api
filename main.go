package main

import (
	"minesweeper/database"
	"minesweeper/database/migrations"
	"minesweeper/server"
)

func main() {
	database.StartDB()
	migrations.RunMigrations(database.GetDatabase())

	server := server.NewServer()
	server.Run()
}
