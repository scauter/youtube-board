package main

import (
	"./my"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main program
func main() {
	my.Migrate()
}
