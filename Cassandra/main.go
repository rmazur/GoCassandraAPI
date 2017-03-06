package Cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("172.20.0.2")
	cluster.Port = 9042
	cluster.Keyspace = "microservpoc"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
