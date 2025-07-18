
package config



import (

	"time"
)

type DBConfig struct {
	Host     string
	Port     int
}



var DefaultConfig = DBConfig{
	Host: "localhost",
	Port:  6379,
}


type SessionData struct {
	ChatID      string
	User        string
	LastSeen    time.Time
	WSConnected int
	Notify      int
}


