package db

import (
	"golang-intern/assignment-2/ent"
	"log"

	"entgo.io/ent/dialect"
)

var Client *ent.Client

func ConnectDb() {
	dbcon, err := ent.Open(dialect.Postgres, "postgresql://postgres:postgres@localhost:5432/earthquake_db?sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	Client = dbcon
	log.Println("Connect")

	/*
		defer dbcon.Close()

		// Run the auto migration tool.
		if err := dbcon.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	*/

}
