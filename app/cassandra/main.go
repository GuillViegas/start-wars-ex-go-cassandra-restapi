package cassandra

import (
	"log"

	"github.com/gocql/gocql"

	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/config"
)

func InitCluster(config *config.CassandraCfg) *gocql.Session {
	cluster := gocql.NewCluster(config.Host)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Username, Password: config.Password}
	cluster.Keyspace = config.Keyspace

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal("FATAL", err)
	}

	return session
}
