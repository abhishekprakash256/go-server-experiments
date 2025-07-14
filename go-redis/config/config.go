
package config

type DBConfig struct {
	Host     string
	Port     int
}



var DefaultConfig = DBConfig{
	Host: "localhost",
	Port:  6379,
}


