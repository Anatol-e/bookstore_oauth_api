package cassandra

import (
	"github.com/gocql/gocql"
	"log"
)

func init() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}
	log.Println("Cassandra db loaded successfully")
	defer session.Close()
}

func RunCassandra() {
	log.Println("logged")
}
