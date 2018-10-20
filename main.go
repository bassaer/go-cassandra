package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

func main() {
	cluster := gocql.NewCluster("192.168.33.10")
	cluster.Keyspace = "mykeyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	var user_id int
	var fname string
	var lname string

	if session == nil {
		log.Fatal("Session closed")
	}

	iter := session.Query("select * from users").Iter()
	for iter.Scan(&user_id, &fname, &lname) {
		fmt.Println(user_id, fname, lname)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
