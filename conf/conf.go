package conf

type Config struct {
	App   *app   `toml:"app"`
	MySQL *mysql `toml:"mysql"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		MySQL: NewDefaultMySQL(),
	}
}

type mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type app struct {
	Name string `toml:"name"`
	Http *http  `toml:"http"`
}

type http struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func NewDefaultMySQL() *mysql {
	return &mysql{
		Host:     "localhost",
		Port:     "3306",
		Database: "vblog",
		Username: "root",
		Password: "123456",
	}
}

func NewDefaultApp() *app {
	return &app{
		Name: "vblog",
		Http: NewDefaultHttp(),
	}
}

func NewDefaultHttp() *http {
	return &http{
		Host: "localhost",
		Port: "8083",
	}
}
