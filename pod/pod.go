package pod

import "github.com/supreme-potato/db"

func Run() {
	// ensure the tables are created
	db.DB().CreateTable("Podspecs")
}
