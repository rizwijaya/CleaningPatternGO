package config

type Conf struct {
	App struct {
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Mode       string `env:"APP_MODE"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET"`
	}
	Db struct {
		Host string `env:"DB_HOST"`
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASSWORD"`
		Port string `env:"DB_PORT"`
	}
	BasicAuth struct {
		Username string `env:"BASIC_AUTH_USER"`
		Password string `env:"BASIC_AUTH_PASSWORD"`
	}
}
