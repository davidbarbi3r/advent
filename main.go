package main

import (
	"example.com/advent/config"
	"example.com/advent/db"
)

func main() {
	config.LoadEnv()
	db.Connect()
}
