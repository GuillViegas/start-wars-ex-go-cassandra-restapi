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
			Host; "",
			Username: "", 
			Password: "",
			Keyspace: "",
		},
	}
}