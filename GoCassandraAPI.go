/*
Simple Go REST API using Cassandra to store data
Date: 22/02/17
Author: Regis Mazur
Email: regis_mazur@hotmail.com
*/

package main

/*
echo "CREATE KEYSPACE microservpoc WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};" | cqlsh
echo "use microservpoc; drop table users; CREATE TABLE users ( id UUID, firstname text, lastname text, age int, email text, city text, PRIMARY KEY (id));" | cqlsh
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rmazur/GoCassandraAPI/Cassandra"
	"github.com/rmazur/GoCassandraAPI/Users"

	"github.com/gorilla/mux"
)

func main() {

	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)

	router.HandleFunc("/users", Users.Get)
	router.HandleFunc("/users/new", Users.Post)
	router.HandleFunc("/users/{user_uuid}", Users.GetOne)

	log.Fatal(http.ListenAndServe(":8080", router))
}

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
