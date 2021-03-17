package config

type Config struct {
	Cassandra *CassandraCfg
}

type CassandraCfg struct {
	Host     string
	Username string
	Password string
	Keyspace string
}

func GetConfig() *Config {
	return &Config{
		Cassandra: &CassandraCfg{
			Host:     "127.0.0.1",
			Username: "planet-app",
			Password: "cGxhbmV0LWFwcA==",
			Keyspace: "star-wars",
		},
	}
}
