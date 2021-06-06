package config

import (
	"flag"
)

var NewFlags struct {
	Port       string
	Mongo      string
	DB         string
	Collection string
}

func SetFlags() {

	flag.StringVar(&NewFlags.Port, "port", "8000", "Port number")
	flag.StringVar(&NewFlags.Mongo, "mongo", "mongodb://root:root@localhost:27017", "Mongo URI")
	flag.StringVar(&NewFlags.DB, "db", "go_search", "Database")
	flag.StringVar(&NewFlags.Collection, "collection", "products", "Collection")
	flag.Parse()
}
