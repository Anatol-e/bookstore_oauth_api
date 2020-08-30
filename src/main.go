package main

import (
	"github.com/Anatol-e/bookstore_oauth_api/src/app"
	"github.com/Anatol-e/bookstore_oauth_api/src/clients/cassandra"
)

func main() {
	cassandra.RunCassandra()
	app.StartApplication()
}
