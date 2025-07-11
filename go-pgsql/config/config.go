package config




type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

var DefaultConfig = DBConfig{
	Host:     "localhost",
	User:     "abhi",
	Password: "mysecretpassword",
	DBName:   "test_db",
	Port:     5432,
}