package db

type Config struct {
	host     string
	username string
	password string
	charset  string
	db       string
	port     string
}

func (config *Config) config() {
	config.username = "Ronnie"
	config.password = "Ronnie.827"
	config.charset = "utf8mb4"
	config.host = "47.98.115.197"
	config.db = "python"
}

func (config *Config) dsn() (dsn string) {
	config.config()
	dsn = config.username + ":" + config.password + "@tcp(" + config.host + ")/" + config.db + "?charset=" + config.charset
	return
}
